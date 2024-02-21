# Create & Delete EC2 credential

This example program demonstrates how to manage creating and deleting EC2 credential for a Service User.

The part of deleting a just-created credential is commented.

## Running this example

Running this file will execute the following operations:

1. **Create:** Create is used to create a new EC2 credential. It is implied, that the Service User ID is known.
2. **(Delete):** _(commented by default)_ Delete deletes a just-created credential on a previous step.

You should see an output like the following (with both operations enabled):

```
Step 1: Created credential Secret Key: a1b2c3... AccessKey: 1a2b3c...
Step 2: Deleted the credential
```
