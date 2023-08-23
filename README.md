# API Documentation

Welcome to the API documentation for our service. This API allows you to manage addresses and travel information, as well as user authentication.

# Public Routes

## Register User

Register a new user.  
**Endpoint:** `POST /v1/auth/register`

**Request:**

    {  
    	"full_name": "John Doe",  
    	"email": "johndoe@example.com",  
    	"password": "secretpassword"  
    }  

**Response:**

    {  
    	"message": "You're registered successfully.",
    	"status": "success"  
    }  

## User Login

Authenticate and login user.  
**Endpoint:** `POST /v1/auth/login`

**Request:**

    {  
    	"email": "johndoe@example.com",  
    	"password": "secretpassword"  
    }  

**Response:**

    {  
    	"accessToken": "access_token",  
    	"refreshToken": "refresh_token"  
    }  

## Refresh Token

Refresh user access token.  
**Endpoint:** `POST /v1/auth/refresh`

**Request:**

    {  
    	"refresh_token": "refresh_token_here"  
    }  

**Response:**

    {  
    	"accessToken": "new_access_token",  
    	"refreshToken": "new_refresh_token"  
    }  

# Protected Routes

Note: For protected routes, include the Authorization header with the value Bearer <access_token>.

## Get User Profile

Get user’s profile information.  
**Endpoint:** `GET /v1/me`  
**Header:** `"Authorization": "Bearer access_token"`

**Response:**

    {  
    	"id": "user_id",  
    	"full_name": "John Doe",  
    	"email": "johndoe@example.com",  
    	"role": "user_role",  
    	"created_at": "2023-08-05T12:34:56Z",  
    	"updated_at": "2023-08-05T12:34:56Z"  
    }  

## Create a Travel

Create a new travel entry.  
**Endpoint:** `POST /v1/travels`  
**Header:** `"Authorization": "Bearer access_token"`

**Request:**

    {
      "visit_date": "2023-08-10T00:00:00Z",
      "location": "Destination City",
      "information": "Travel details here", ? Optional
      "image_url": "https://example.com/image.png", ? Optional
      "latitude": 37.12353,
      "longitude": -122.95421
    }

**Response:**

    {
       "message": "Travel successfully created",
       "status": "success"
    }

## List Travels

Get a list of travel entries.

**Endpoint:** `GET /v1/travels?page=1&limit=10`  
**Header:** `"Authorization": "Bearer access_token"`

**Response:**

    {
      "data": {
        "count": 3,
        "travels": [
          {
            "id": "39fa7def-a4ce-42e1-8f66-97b6fdecd639",
            "visit_date": "2023-08-10T00:00:00Z",
            "location": "Destination City",
            "information": "Travel details here",
            "image_url": "URL here",
            "latitude": 42.0,
            "longitude": 122.0,
            "created_at": "2023-08-18T20:49:46.869205Z",
            "updated_at": "2023-08-18T20:49:46.869205Z"
          },
          {
            "id": "8d76b17b-67fa-4afe-a335-b0c089e50255",
            "visit_date": "2023-08-10T00:00:00Z",
            "location": "Destination City",
            "information": "Travel details here",
            "image_url": "URL here",
            "latitude": 42.0,
            "longitude": 122.0,
            "created_at": "2023-08-18T20:49:47.678662Z",
            "updated_at": "2023-08-18T20:49:47.678662Z"
          },
          {
            "id": "01abd8f0-2e1d-4a39-9e8a-7fb55da83760",
            "visit_date": "2023-08-10T00:00:00Z",
            "location": "Destination City",
            "information": "Travel details here",
            "image_url": "URL here",
            "latitude": 42.0,
            "longitude": 122.0,
            "created_at": "2023-08-18T20:49:51.010944Z",
            "updated_at": "2023-08-18T20:49:51.010944Z"
          }
        ]
      },
      "status": "success"
    }

## Get Travel by ID

Get details of a specific travel entry.

**Endpoint:** `GET /v1/travels/travelId`  
**Header:** `"Authorization": "Bearer access_token"`  
**Response:**

    {
    	"data": {
    		"travel": {
    			"id": "0d498d8a-2c2a-4908-b03b-e153762cecde",
    			"visit_date": "2023-08-10T00:00:00Z",
    			"location": "Destination City",
    			"information": "Travel details here",
    			"image_url": "https://example.com/image2.jpg",
    			"latitude": 48.8566,
    			"longitude": 122.12353,
    			"created_at": "2023-08-18T20:49:40.687527Z",
    			"updated_at": "2023-08-18T20:49:40.687527Z"
    		}
    	},
    	"status": "success"
    }

## Update Travel

Update details of a specific travel entry.

**Endpoint:** `PUT /v1/travels/travelId`  
**Header:** `"Authorization": "Bearer access_token"`  
**Request:**

    {
      "visit_date": "2023-08-10T00:00:00Z",
      "location": "Destination City",
      "information": "Travel details here", ? Optional
      "image_url": "https://example.com/image.png", ? Optional
      "latitude": 37.12353,
      "longitude": -122.95421
    }

**Response:**

    {
    	"status": "success",
    	"message": "Travel successfully updated"
    }

## Delete Travel

Delete a specific travel entry.

**Endpoint:** `DELETE /v1/travels/travelId`  
**Header:** `"Authorization": "Bearer access_token"`  
**Response:**

    {
    	"message": "travel successfully deleted",
    	"status": "success"
    }

## Create a Gallery Image

Create a new gallery entry.  
**Endpoint:** `POST /v1/galleries`  
**Header:** `"Authorization": "Bearer access_token"`

**Request:**

    {
      "travel_id": "358c3c03-a66a-4d03-adc7-84a1d9874d6e",
      "image_url": "https://example.com/tree.png",
      "caption": "Bir ağaç resmi" 
    }

**Response:**

    {
    	"message": "Image added to gallery",
    	"status": "success"
    }

## Get All Gallery Images

Get all gallery images.  
**Endpoint:** `GET /v1/galleries/travelId`  
**Header:** `"Authorization": "Bearer access_token"`

**Response:**

    {
    	"data": {
    		"images": [
    			{
    				"id": "39e25b3e-fe18-45d1-98c7-4c22c56c0d33",
    				"travel_id": "358c3c03-a66a-4d03-adc7-84a1d9874d6e",
    				"image_url": "https://example.org/tree.png",
    				"caption": "Bir ağaç resmi",
    				"created_at": "2023-08-18T21:06:26.866468Z",
    				"updated_at": "2023-08-18T21:06:26.866468Z"
    			}
    		],
    		"count": 1
    	},
    	"status": "success"
    }

## Delete Gallery Image

Delete a specific gallery image.  
**Endpoint:** `DELETE /v1/galleries/travelId/imageId`  
**Header:** `"Authorization": "Bearer access_token"`

**Response:**

    {
    	"message": "Image deleted from gallery",
    	"status": "success"
    }

### Address

Perform operations related to addresses.

## Create Address

Create a new address entry.

**Endpoint:** `POST /v1/addresses`  
**Header:** `"Authorization": "Bearer access_token"`  
**Request:**

    {  
    	"address_title": "Home",  
    	"state": "State",  
    	"city": "City",  
    	"country": "Country",  
    	"address": "Address"  
    }  

**Response:**

    {
    	"message": "Address successfully created",
    	"status": "success"
    }

## List Addresses

Get a list of address entries.

**Endpoint:** `GET /v1/addresses`  
**Header:** `"Authorization": "Bearer access_token"`  
**Response:**

    {
    	"data": {
    		"addresses": [
    			{
    				"id": "39e25b3e-fe18-45d1-98c7-4c22c56c0d33",
    				"address_title": "Home",
    				"state": "State",
    				"city": "City",
    				"country": "Country",
    				"address": "Address",
    				"created_at": "2023-08-18T21:06:26.866468Z",
    				"updated_at": "2023-08-18T21:06:26.866468Z"
    			}
    		],
    		"count": 1
    	},
    	"status": "success"
    }

## Get an Address by ID

Get details of a specific address entry.

**Endpoint:** `GET /v1/addresses/$(addressId)`  
**Header:** `"Authorization": "Bearer access_token"`  
**Response:**

    {
    	"data": {
    		"address": {
    			"id": "39e25b3e-fe18-45d1-98c7-4c22c56c0d33",
    			"address_title": "Home",
    			"state": "State",
    			"city": "City",
    			"country": "Country",
    			"address": "Address",
    			"created_at": "2023-08-18T21:06:26.866468Z",
    			"updated_at": "2023-08-18T21:06:26.866468Z"
    		}
    	},
    	"status": "success"
    }

## Update Address

Update details of a specific address entry.

**Endpoint:** `PUT /v1/addresses/$(addressId)`  
**Header:** `"Authorization": "Bearer access_token"`  
**Request:**

    {  
    	"address_title": "Home",  
    	"state": "State",  
    	"city": "City",  
    	"country": "Country",  
    	"address": "Address"  
    }  

**Response:**

    {
    	"status": "success",
    	"message": "Address successfully updated"
    }

## Delete Address

Delete a specific address entry.

**Endpoint:** `DELETE /v1/addresses/$(addressId)`  
**Header:** `"Authorization": "Bearer access_token"`  
**Response:**

    {
    	"message": "address successfully deleted",
    	"status": "success"
    }

</div>

</div>
