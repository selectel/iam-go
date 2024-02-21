package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client/testdata"
)

func TestDoRequest(t *testing.T) {
	type args struct {
		body   io.Reader
		method string
		url    string
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedBody  []byte
		expectedError error
	}{
		{
			name: "Test DoRequest GET method",
			args: args{
				method: http.MethodGet,
				url:    testdata.TestURL,
				body:   nil,
			},
			prepare: func() {
				httpmock.RegisterResponder(http.MethodGet, testdata.TestURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(200, testdata.TestDoRequestRaw)
						return resp, nil
					})
			},
			expectedBody:  []byte(testdata.TestDoRequestRaw),
			expectedError: nil,
		},
		{
			name: "Test DoRequest POST method",
			args: args{
				method: http.MethodPost,
				url:    testdata.TestURL,
				body:   bytes.NewReader([]byte(testdata.TestDoRequestRaw)),
			},
			prepare: func() {
				httpmock.RegisterResponder(http.MethodPost, testdata.TestURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(200, testdata.TestDoRequestRaw)
						return resp, nil
					})
			},
			expectedBody:  []byte(testdata.TestDoRequestRaw),
			expectedError: nil,
		},
		{
			name: "Test DoRequest GET method return Error",
			args: args{
				method: http.MethodGet,
				url:    testdata.TestURL,
				body:   nil,
			},
			prepare: func() {
				httpmock.RegisterResponder(http.MethodGet, testdata.TestURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(403, testdata.TestDoRequestErr)
						return resp, nil
					})
			},
			expectedBody:  nil,
			expectedError: iamerrors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			baseClient := &BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     tt.args.url,
				AuthMethod: &KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
				UserAgent:  testdata.TestUserAgent,
			}

			httpmock.ActivateNonDefault(baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actualBody, err := baseClient.DoRequest(ctx, DoRequestInput{
				Body:   tt.args.body,
				Method: tt.args.method,
				Path:   "/",
			})

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedBody, actualBody)
		})
	}
}
