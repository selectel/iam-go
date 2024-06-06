# Create Group with Role & add User

This example program demonstrates how to manage creating and deleting Group with Roles and Users.

The part of deleting is disabled by `deleteAfterRun` variable.

As an example, the Member Role will be assigned for a new Group.

## Running this example

Running this file will execute the following operations:

1. **Create Group:** Create is used to create a new Group.
2. **Create User:** Create is used to create a new User.
3. **Assign Role:** Assign a role to the created Group.
4. **Update Group:** Updates the Group Name and Description.
5. **(Delete Group):** _(disabled by default)_ Delete a just-created Group on a previous step.
6. **(Delete User):** _(disabled by default)_ Delete a just-created User on a previous step.

You should see an output like the following:
```
Step 1: Created Group Name: test_group_name ID: 1a2b3c...
Step 2: Created User ID: 12345_3... Keystone ID: 1a2b3c...
Step 3: Assigned Role member with scope account to Group ID: 1a2b3c...
Step 4: Group Name and Description updated to: new_test_group_name and new_group_description
Step 5: Deleting Group with ID: 1a2b3c...
Step 6: Deleting User with ID: 12345_3...
```
