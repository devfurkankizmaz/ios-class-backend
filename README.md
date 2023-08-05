# API Documentation

Welcome to the API documentation for our service. This API allows you to manage addresses and travel information, as well as user authentication.

## Public Routes

### Register User
Register a new user.
**Endpoint:** `POST /v1/auth/register`

**Request:**
```json
{
  "full_name": "John Doe",
  "email": "johndoe@example.com",
  "password": "secretpassword"
}
```
**Response:**
```json
{
  "message": "Inserted ID: d00e6165-8453-4e69-b07d-82a87cfb51fa",
  "status": "success"
}
```

### Refresh Token
Refresh user access token.
**Endpoint:** `POST /v1/auth/refresh`

**Request:**
```json
{
  "refresh_token": "refresh_token_here"
}

```
**Response:**
```json
{
  "accessToken": "new_access_token",
  "refreshToken": "new_refresh_token"
}
```

### User Login
Authenticate and login user.
**Endpoint:** `POST /v1/auth/login`

**Request:**
```json
{
  "email": "johndoe@example.com",
  "password": "secretpassword"
}
```
**Response:**
```json
{
  "accessToken": "access_token",
  "refreshToken": "refresh_token"
}
```

## Protected Routes
Note: For protected routes, include the Authorization header with the value Bearer <access_token>.

### Get User Profile
Get user's profile information.
**Endpoint:** `GET /v1/me`

**Response:**
```json
{
  "id": "user_id",
  "full_name": "John Doe",
  "email": "johndoe@example.com",
  "role": "user_role",
  "created_at": "2023-08-05T12:34:56Z",
  "updated_at": "2023-08-05T12:34:56Z"
}
```

### Create a Travel
Create a new travel entry.
**Endpoint:** `POST /v1/travels`

**Request:**
```json
{
  "visit_date": "2023-08-10T00:00:00Z",
  "location": "Destination City",
  "information": "Travel details here"
}
```
**Response:**
```json
{
  "visit_date": "2023-08-10T00:00:00Z",
  "location": "Destination City",
  "information": "Travel details here",
  "created_at": "2023-08-05T12:34:56Z",
  "updated_at": "2023-08-05T12:34:56Z"
}
```

### List Travels

Get a list of travel entries.

**Endpoint:** `GET /v1/travels?page=1&limit=10`

### Get Travel by ID

Get details of a specific travel entry.

**Endpoint:** `GET /v1/travels/$(travelId)`

### Update Travel

Update details of a specific travel entry.

**Endpoint:** `PUT /v1/travels/$(travelId)`

### Delete Travel

Delete a specific travel entry.

**Endpoint:** `DELETE /v1/travels/$(travelId)`

## Address

Perform operations related to addresses.

### Create Address

Create a new address entry.

**Endpoint:** `POST /v1/addresses`
**Request:**
```json
{
  "address_title": "Home",
  "state": "State",
  "city": "City",
  "country": "Country",
  "address": "Address"
}
```
**Response:**
```json
{
  "address_title": "Home",
  "state": "State",
  "city": "City",
  "country": "Country",
  "address": "Address",
  "created_at": "2023-08-05T12:34:56Z",
  "updated_at": "2023-08-05T12:34:56Z"
}
```

### List Addresses

Get a list of address entries.

**Endpoint:** `GET /v1/addresses`

### Get Address by ID

Get details of a specific address entry.

**Endpoint:** `GET /v1/addresses/$(addressId)`

### Update Address

Update details of a specific address entry.

**Endpoint:** `PUT /v1/addresses/$(addressId)`

### Delete Address

Delete a specific address entry.

**Endpoint:** `DELETE /v1/addresses/$(addressId)`

