# Create, Update & Delete Service User

This example program demonstrates how to manage creating, updating and deleting Service User.

The part of deleting a just-created Service User is commented.

As an example, the Billing Role will be assigned for a new Service User and in update method this Service User will be set to _Disabled_.

## Running this example

Running this file will execute the following operations:

1. **Create:** Create is used to create a new Service User.
2. **Update** Update sets _Enabled_ property of the just-created Service User to _false_
3. **(Delete):** _(commented by default)_ Delete deletes a just-created Service User.

You should see an output like the following (with all operations enabled):

```
Step 1: Created Service User ID: a1b2c3...
Step 2: Disabled Service User ID: a1b2c3...
Step 3: Deleted Service User ID: a1b2c3...
```
