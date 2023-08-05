# Auth Endpoint Documentation

## Login

**Request**
```json
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "secretpassword"
}
```
**Response**
```json
Status: 200 OK
Content-Type: application/json

{
  "accessToken": "your-access-token",
  "refreshToken": "your-refresh-token"
}
```

## Register
***Request***
```json
POST /api/auth/register
Content-Type: application/json

{
"full_name": "John Doe",
"email": "user@example.com",
"password": "secretpassword"
}

```
***Response***
```json
Status: 201 Created
Content-Type: application/json

{
"message": "Inserted ID: UUID:V4",
}
```

## Refresh
***Request***
```json
POST /api/auth/refresh
Content-Type: application/json

{
"refresh_token": "your-refresh-token"
}

```
***Response***
```json
Status: 200 OK
Content-Type: application/json

{
"accessToken": "your-new-access-token",
"refreshToken": "your-new-refresh-token"
}
```

## Fetch Profile
***Request***
```json
GET /api/me
Authorization: Bearer your-access-token

```
***Response***
```json
Status: 200 OK
Content-Type: application/json

{
"id": "user-id",
"full_name": "John Doe",
"email": "user@example.com"
}

```

# Travel Endpoint Documentation

## Create a New Travel

**Request**
```json
POST /api/travels
Content-Type: application/json
Authorization: Bearer your-access-token

{
  "visit_date": "2023-07-30T12:00:00Z",
  "location": "New York",
  "information": "A trip to the Big Apple!"
}
```
**Response**
```json

Status: 200 OK
Content-Type: application/json

{
"id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
"user_id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
"visit_date": "2023-07-30T12:00:00Z",
"location": "New York",
"information": "A trip to the Big Apple!",
"created_at": "2023-07-30T10:30:00Z",
"updated_at": "2023-07-30T10:30:00Z"
}
```
## Fetch All Travels by User ID
**Request**
```json
GET /api/travels?page=1&limit=10
Authorization: Bearer your-access-token
```
**Response**
```json
Status: 200 OK
Content-Type: application/json

[
  {
    "id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
    "user_id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
    "visit_date": "2023-07-30T12:00:00Z",
    "location": "New York",
    "information": "A trip to the Big Apple!",
    "created_at": "2023-07-30T10:30:00Z",
    "updated_at": "2023-07-30T10:30:00Z"
  },
  {
    "id": "b5a7e6b7-443b-4b64-9a3f-3c4c6d778899",
    "user_id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
    "visit_date": "2023-08-10T15:00:00Z",
    "location": "London",
    "information": "Exploring the UK!",
    "created_at": "2023-07-31T08:45:00Z",
    "updated_at": "2023-07-31T08:45:00Z"
  }
]
```

## Fetch By ID
**Request**
```json
GET /api/travels/:travelId
Authorization: Bearer your-access-token
```
**Response**
```json
Status: 200 OK
Content-Type: application/json

{
  "id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
  "user_id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
  "visit_date": "2023-07-30T12:00:00Z",
  "location": "New York",
  "information": "A trip to the Big Apple!",
  "created_at": "2023-07-30T10:30:00Z",
  "updated_at": "2023-07-30T10:30:00Z"
}
```
## Delete Travel by ID
**Request**
```json
DELETE /api/travels/:travelId
Authorization: Bearer your-access-token
```

# Address Endpoint Documentation

## Create a New Address

**Request**
```json
POST /api/addresses
Content-Type: application/json
Authorization: Bearer your-access-token

{
    "address_title": "Home",
    "state": "California",
    "city": "Los Angeles",
    "country": "USA",
    "address": "123 Main Street"
}
```
**Response**
```json

Status: 200 OK
Content-Type: application/json

{
    "address_title": "Home",
    "state": "California",
    "city": "Los Angeles",
    "country": "USA",
    "address": "123 Main Street",
    "created_at": "2023-08-05T12:34:56Z",
    "updated_at": "2023-08-05T12:34:56Z"
}
```

## Delete an Address

```json
DELETE /api/addresses/:addressId
Authorization: Bearer your-access-token
```
## Fetch Address
**Request**
```json
GET /api/travels/:travelId
Authorization: Bearer your-access-token
```
**Response**
```json
Status: 200 OK
Content-Type: application/json

{
    "address_title": "Home",
    "state": "California",
    "city": "Los Angeles",
    "country": "USA",
    "address": "123 Main Street",
    "created_at": "2023-08-05T12:34:56Z",
    "updated_at": "2023-08-05T12:34:56Z"
}
```

## Fetch All Addresses
**Request**
```json
GET /api/addresses
Authorization: Bearer your-access-token
```
**Response**
```json
Status: 200 OK
Content-Type: application/json

[
  {
    "id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
    "user_id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
    "address_title": "Home",
    "state": "California",
    "city": "Los Angeles",
    "country": "USA",
    "address": "123 Main Street",
    "created_at": "2023-08-05T12:34:56Z",
    "updated_at": "2023-08-05T12:34:56Z"
  },
  {
    "id": "b5a7e6b7-443b-4b64-9a3f-3c4c6d778899",
    "user_id": "a6b7d7c6-344a-4a85-8e8e-6e1b1b234567",
    "address_title": "Home",
    "state": "California",
    "city": "Los Angeles",
    "country": "USA",
    "address": "123 Main Street",
    "created_at": "2023-08-05T12:34:56Z",
    "updated_at": "2023-08-05T12:34:56Z"
  }
]
```
