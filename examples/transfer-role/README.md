# Transfer role from one User to another

This example program demonstrates how to unassign Billing role from one User and assign it to another.

The same approach can be applied for Service Users.

## Running this example

Running this file will execute the following operations:

1. **List:** List is used to retrieve all Users. The first one, who has a billing role, will be selected as 'transferer'.
2. **UnassignRole:** UnassinRole will remove Billing role from chosen user.
3. **AssignRole:** AssignRole will add Billing role to the predefined User ID.

You should see an output like the following:

```
Step 1: User 123456_12345 with the Billing role was found
Step 2: Unassigned the Billing role from User 123456_12345 
Step 3: Assigned the Billing role to User 654321_65432 
```
