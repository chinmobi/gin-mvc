## The RESTful APIs

### NOTE:

Each of the APIs will return the unique error response format if some errors occurred:
<br/>(The field like `"_comment": "..."`, is not the real field, just as a comment)

```json
{
  "apiVersion": "v1",
  "error": {
    "statusCode": "<the status code>",
    "code": "<the error code identifying the specific error>",
    "message": "<the brief description about the error>",
    "errors": [ {"_comment": "Array of wrapped errors' details."},
      {
        "name": "<error type name>",
        "message": "<error message>",
        "_other_error_fields": "..."
      },
      {"_other_errors": "..."}
    ]
  }
}
```

For normal response, the response's data entity format is like:

```json
{
  "apiVersion": "v1",
  "data": { "_comment": "Any object.",
    "_each_of_the_data_fields": "..."
  }
}
```
or array of objects:

```json
{
  "apiVersion": "v1",
  "data": [ {"_comment": "Array of objects."},
    {"_each_of_the_objects": "..."}
  ]
}
```

### DEMO for: Create, Read, Update and Delete `User`

* Get all users

```
GET /api/v1/users
```

Normal response body:

```json
{
  "apiVersion": "v1",
  "data": [ {"_comment": "Array of user objects"},
    { "_comment": "Each of the user object",
      "_each_of_the_user_fields": "..."
    },
    {"_other_user_objects": "..."}
  ]
}
```

* Get a user by user's id (uid)

```
GET /api/v1/users/:uid
```

Normal response body:

```json
{
  "apiVersion": "v1",
  "data": { "_comment": "The user object",
    "id": "<the user's id>",
    "nickname": "<the user's nickname>",
    "email": "<the user's email>"
  }
}
```

* Create a user

```
POST /api/v1/users
```

Request body:

```json
{
  "nickname": "<the user's nickname>",
  "email": "<the user's email>"
}
```

Normal response body:

```json
{
  "apiVersion": "v1",
  "data": { "_comment": "The user object",
    "id": "<the user's id>",
    "nickname": "<the user's nickname>",
    "email": "<the user's email>"
  }
}
```

* Update the user

```
PUT /api/v1/users/:uid
```
or

```
PATCH /api/v1/users/:uid
```

Request body:

```json
{
  "nickname": "<the user's nickname>",
  "email": "<the user's email>"
}
```

Normal response body:

```json
{
  "apiVersion": "v1",
  "data": { "_comment": "The user object",
    "id": "<the user's id>",
    "nickname": "<the user's nickname>",
    "email": "<the user's email>"
  }
}
```

* Delete the user

```
DELETE /api/v1/users/:uid
```

None response content.

* Demo for more complex RESTful APIs

```
// Get the user(identified by the uid)'s all roles:
GET /api/v1/users/:uid/roles

// Get the user(identified by the uid)'s role by role's id (rid):
GET /api/v1/users/:uid/roles/:rid

// Add a role for the user(identified by the uid)
POST /api/v1/users/:uid/roles

// Update the user(identified by the uid)'s role(identified by the rid):
PUT /api/v1/users/:uid/roles/:rid

// Delete the user(identified by the uid)'s role(identified by the rid):
DELETE /api/v1/users/:uid/roles/:rid

```

### DEMO for: Login, Logout, Register `User account`

* Login the user's account

```
POST /login
```

Request body:

```json
{
  "username": "<the user's nickname or email>",
  "password": "<the user's account password>"
}
```

Normal response body:

```json
{
  "apiVersion": "v1",
  "data": { "_comment": "The user object",
    "id": "<the user's id>",
    "nickname": "<the user's nickname>",
    "email": "<the user's email>"
  }
}
```

* Register a user's account

```
POST /signup
```

Request body:

```json
{
  "nickname": "<the user's nickname>",
  "password": "<the user's account password>",
  "email": "<the user's email>"
}
```

Normal response body:

```json
{
  "apiVersion": "v1",
  "data": { "_comment": "The user object",
    "id": "<the user's id>",
    "nickname": "<the user's nickname>",
    "email": "<the user's email>"
  }
}
```

* Logout the user's account

```
POST /auth/logout?uid=<user's id>
```
or

```
GET /auth/logout?uid=<user's id>
```

The response should be redirected to the `/login` URL.

* The `/auth/**` path could be used for more account actions, such as:

```
// Reset the account's password
POST /auth/reset_passwd

// Refresh the account's auth token
POST /auth/refresh_token

```

### Best Practices

* The `/public/**` should be used for static resources (*.html, *.css, *.js, *.jpg, ...), NOT be authenticated.

* The `/api/v1/**` used for RESTful APIs resources, SHOULD be authenticated.

* The `/login` used for login account, NOT be authenticated.

* The `/signup` used for register an account, NOT be authenticated.

* The `/auth/**` used for account's auth actions, MUST be authenticated.

* Using `/api/v1/**`, `/api/v2/**`, ... to identify api versions.
