# Error Handling

In generall, iam-go errors can be divided into two types:

1. **IAM Server Errors**: returned by the IAM API Server itself
2. **IAM Client Errors**: returned by the iam-go library

Any of these can be handled as a special type: [_**iamerrors.Error**_](../iamerrors/iamerrors.go). 

Below are some examples on how to handle these errors:

1. You can use _errors.Is_ to identify the type of error:

```go
if err != nil {
    switch {
    case errors.Is(err, iamerrors.ErrForbidden):
    	log.Fatalf("No rights: %s", err.Error())
    }
    ...
}
```

2. You can cast a returned error to _iamerrors.Error_ with _errors.As_ and get the specific info (description of an error, for example):

```go
if err != nil {
    var iamError *iamerrors.Error
    if errors.As(err, &iamError) {
        log.Fatalf("IAM Error! Description: %s", iamError.Desc)
    }
    ...
}
```