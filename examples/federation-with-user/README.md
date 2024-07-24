# Create Federation with Certificate & add User

This example program demonstrates how to manage creating and deleting Federation with Certificate and assigning Users.

The part of deleting is disabled by `deleteAfterRun` variable.

## Running this example

Running this file will execute the following operations:

1. **Create Federation:** Create is used to create a new Federation.
2. **Create Certificate for Federation:** Create is used to create a new Certificate for Federation.
3. **Create federated User:** Create is used to create a new federated User.
4. **Update Federation:** Updates the Federation Name and Description.
5. **(Delete Federation):** _(disabled by default)_ Delete a just-created Federation on a previous step.
6. **(Delete Federation Certificate):** _(disabled by default)_ Delete a just-created Federation Certificate on a previous step.
7. **(Delete User):** _(disabled by default)_ Delete a just-created federated User on a previous step.

You should see an output like the following:
```
Step 1: Created Federation Name: federation_name ID: 1a2b3c...
Step 2: Created Certificate for Federation ID: 12345_3... Federation ID: 1a2b3c...
Step 3: Created federated User ID: 54321_2... Keystone ID: 1c2b3a...
Step 4: Updated Federation Name and Description
Step 5: Deleting Federation with ID: 1a2b3c...
Step 6: Deleting Federation Certificate with ID: 12345_3...
Step 6: Deleting User with ID: 54321_2...
```
