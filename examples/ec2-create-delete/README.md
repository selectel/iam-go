# Create & Delete S3 Credentials

This example program demonstrates how to manage creating and deleting S3 Credentials for a Service User.

The part of deleting a just-created credentials is commented.

## Running this example

Running this file will execute the following operations:

1. **Create:** Create is used to create a new S3 Credentials. It is implied, that the Service User ID is known.
2. **(Delete):** _(commented by default)_ Delete deletes a just-created credentials on a previous step.

You should see an output like the following (with both operations enabled):

```
Step 1: Created credentials Secret Key: a1b2c3... AccessKey: 1a2b3c...
Step 2: Deleted the credentials
```
