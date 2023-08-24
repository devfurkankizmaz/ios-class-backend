# Travio API Documentation

Welcome to the API documentation for **Travio App** Service. This API allows you to manage places, galleries, visits and addresses information, as well as user authentication.

# Auth & Upload
## Sign Up
Register New User
**Endpoint:** `POST /v1/auth/register`

**Request**
```  
{  
	"full_name": "John Doe",  
	"email": "johndoe@example.com",  
	"password": "secretpassword"  
}
```
**Response**
```  
{  
	"message": "You're registered successfully.",  
	"status": "success"  
}
```
## Login
Authenticate and login user.
**Endpoint:** `POST /v1/auth/login`

**Request**
```  
{
	"email":  "johndoe@example.com",
	"password":  "secretpassword"  
}
```
**Response**
```  
{
	"accessToken":  "access_token",
	"refreshToken":  "refresh_token"  
}
```
## Refresh
Refresh user access token.
**Endpoint:** `POST /v1/auth/refresh`

**Request**
```  
{
	"refresh_token":  "refresh_token_here"  
}
```
**Response**
```
{
	"accessToken":  "access_token",
	"refreshToken":  "refresh_token"  
}
```
## Upload
Upload image with multipart form data.
**Endpoint:** `POST /upload`
**Content-Type** `multipart/form-data`

**Request Body**
key: "file": The image file to be uploaded (allowed extensions: `.jpg`, `.jpeg`, `.png`)

**Response**
```
{
	"messageType": "S",
	"message": "Files uploaded successfully",
	"urls": [
	 "https://iosclass.ams3.digitaloceanspaces.com/1631234567890.jpg",
	 "https://iosclass.ams3.digitaloceanspaces.com/1631234567891.jpg" 
   ] 
 }
```
## Profile
Get user profile by auth header
**Endpoint:** `GET v1/me`
**Header** `"Authorization": "Bearer access_token"`

**Response**
```
{
	"id":  "user_id",
    "full_name":  "John Doe",
    "email":  "johndoe@example.com",
    "role":  "user_role",
    "created_at":  "2023-08-05T12:34:56Z",
    "updated_at":  "2023-08-05T12:34:56Z"  
}
```
# Places
## Post a place
Create a place
**Endpoint:** `POST /v1/places`
**Header** `"Authorization": "Bearer access_token"`

**Request**
```  
{
	"place": "Nevşehir, Türkiye",
    "title": "Kapadokya",
    "description": "Kapadokya, adeta peri masallarının gerçeğe dönüştüğü büyülü bir dünyadır.",
    "cover_image_url": "https://iosclass.ams3.digitaloceanspaces.com/1692817007606598873.png",
    "latitude": 38.6431,
    "longitude": 34.8287
}
```
**Response**
```  
{
	"message": "Place successfully created.",
	"status": "success"
}
```
## Update a place
Update a place
**Endpoint:** `PUT /v1/places`
**Header** `"Authorization": "Bearer access_token"`

**Request**
```  
{
	"place": "Nevşehir, Türkiye",
    "title": "Kapadokya",
    "description": "Kapadokya, adeta peri masallarının gerçeğe dönüştüğü büyülü bir dünyadır.",
    "cover_image_url": "https://iosclass.ams3.digitaloceanspaces.com/1692817007606598873.png",
    "latitude": 38.6431,
    "longitude": 34.8287
}
```
**Response**
```  
{
	"message": "Place successfully updated.",
	"status": "success"
}
```
## Delete a place
Delete a place
**Endpoint:** `DELETE /v1/places/:placeId`
**Header** `"Authorization": "Bearer access_token"`

**Response**
```  
{
	"message": "Place successfully deleted.",
	"status": "success"
}
```
## Get All Places
Get all places
**Endpoint:** `GET /v1/places`
**Pagination:** `GET /v1/places?page=1&limit=10`

**Response**
```  
{
  "data": {
    "count": 1,
    "places": [
      {
        "id": "2983980a-4035-4211-a98d-deb264c6f9f5",
        "creator": "Furkan Kızmaz",
        "place": "Nevşehir, Türkiye",
        "title": "Kapadokya",
        "description": "Description...",
        "cover_image_url": "https://iosclass.ams3.digitaloceanspaces.com/1692817007606598873.png",
        "latitude": 38.6454,
        "longitude": 34.8283,
        "created_at": "2023-08-24T08:39:21.978094Z",
        "updated_at": "2023-08-24T08:42:48.150143Z"
      }
    ]
  },
  "status": "success"
}
```
## Get a Place by ID
Get all places
**Endpoint:** `GET /v1/places/:placeId`

**Response**
```
{
  "data": {
    "place": {
      "id": "5ee3e518-39e5-47b6-a18a-e942a5598aae",
      "creator": "Melihozm",
      "place": "Adıyaman, Türkiye",
      "title": "Nemrut Dağı Milli Parkı",
      "description": "Description...",
      "cover_image_url": "https://live.staticflickr.com/2941/15102053140_69e59cc770_b.jpg",
      "latitude": 37.980927,
      "longitude": 38.74131,
      "created_at": "2023-08-24T08:54:08.329671Z",
      "updated_at": "2023-08-24T08:54:08.329671Z"
    }
  },
  "status": "success"
}
```
## Get All Places for User
Get all places specify with user auth token
**Endpoint:** `GET /v1/places/user`
**Header** `"Authorization": "Bearer access_token"`

**Response**
```
{
  "data": {
    "place": {
      "id": "5ee3e518-39e5-47b6-a18a-e942a5598aae",
      "creator": "Melihozm",
      "place": "Adıyaman, Türkiye",
      "title": "Nemrut Dağı Milli Parkı",
      "description": "Description...",
      "cover_image_url": "https://live.staticflickr.com/2941/15102053140_69e59cc770_b.jpg",
      "latitude": 37.980927,
      "longitude": 38.74131,
      "created_at": "2023-08-24T08:54:08.329671Z",
      "updated_at": "2023-08-24T08:54:08.329671Z"
    }
  },
  "status": "success"
}
```
# Gallery
## Post a gallery image
Create a gallery image
**Endpoint:** `POST /v1/galleries`
**Header** `"Authorization": "Bearer access_token"`

**Request**
```  
{
    "place_id": "358c3c03-a66a-4d03-adc7-84a1d9874d6e",
    "image_url": "https://example.com/animage.png"
}
```
**Response**
```  
{
	"message": "Image added to gallery.",
	"status": "success"
}
```
## Delete a gallery image
Delete a gallery image
**Endpoint:** `DELETE /v1/galleries/:placeId/:imageId`
**Header** `"Authorization": "Bearer access_token"`

**Response**
```  
{
	"message": "Image deleted from gallery.",
	"status": "success"
}
```
## Get All Gallery by Place ID
It's public route dont, need to give header.
**Endpoint:** `GET /v1/galleries/:placeId`

**Response**
```
{
  "data": {
    "count": 3,
    "images": [
      {
        "id": "287c0591-f0c2-4e67-9061-6d2d3eba7848",
        "place_id": "2983980a-4035-4211-a98d-deb264c6f9f5",
        "image_url": "https://iosclass.ams3.digitaloceanspaces.com/1692868653636050606.png",
        "created_at": "2023-08-24T09:18:18.759584Z",
        "updated_at": "2023-08-24T09:18:18.759584Z"
      },
      {
        "id": "2309dbe0-b70d-40ae-ba09-2c565bb17993",
        "place_id": "2983980a-4035-4211-a98d-deb264c6f9f5",
        "image_url": "https://iosclass.ams3.digitaloceanspaces.com/1692868654206658092.png",
        "created_at": "2023-08-24T09:18:30.33605Z",
        "updated_at": "2023-08-24T09:18:30.33605Z"
      },
      {
        "id": "e36f8c3d-0052-478c-810b-521336e5aa41",
        "place_id": "2983980a-4035-4211-a98d-deb264c6f9f5",
        "image_url": "https://iosclass.ams3.digitaloceanspaces.com/1692868654322256108.png",
        "created_at": "2023-08-24T09:18:39.709658Z",
        "updated_at": "2023-08-24T09:18:39.709658Z"
      }
    ]
  },
  "status": "success"
}
```
# Visit
## Post a Visit
Create a user visit
**Endpoint:** `POST /v1/visits`
**Header** `"Authorization": "Bearer access_token"`

**Request**
```  
{
    "place_id": "358c3c03-a66a-4d03-adc7-84a1d9874d6e",
    "visited_at": "2023-08-10T00:00:00Z"
}
```
**Response**
```  
{
	"message": "Visit successfully created.",
	"status": "success"
}
```
## Get All Visits
Get all user visits
**Endpoint:** `GET /v1/visits`
**Pagination:** `GET /v1/visits?page=1&limit=10`
**Header** `"Authorization": "Bearer access_token"`

**Response**
```
{
  "data": {
    "count": 3,
    "visits": [
      {
        "id": "287c0591-f0c2-4e67-9061-6d2d3eba7848",
        "place_id": "2983980a-4035-4211-a98d-deb264c6f9f5",
        "visited_at": "2023-08-24T09:18:18.759584Z",
        "created_at": "2023-08-24T09:18:18.759584Z",
        "updated_at": "2023-08-24T09:18:18.759584Z"
      },
      {
        "id": "2309dbe0-b70d-40ae-ba09-2c565bb17993",
        "place_id": "2983980a-4035-4211-a98d-deb264c6f9f5",
        "visited_at": "2023-08-24T09:18:18.759584Z",
        "created_at": "2023-08-24T09:18:30.33605Z",
        "updated_at": "2023-08-24T09:18:30.33605Z"
      },
      {
        "id": "e36f8c3d-0052-478c-810b-521336e5aa41",
        "place_id": "2983980a-4035-4211-a98d-deb264c6f9f5",
        "visited_at": "2023-08-24T09:18:18.759584Z",
        "created_at": "2023-08-24T09:18:39.709658Z",
        "updated_at": "2023-08-24T09:18:39.709658Z"
      }
    ]
  },
  "status": "success"
}
```
## Get A Visit By ID
Get a visit by visit id.
**Endpoint:** `GET /v1/visits/visitId`
**Header** `"Authorization": "Bearer access_token"`

**Response**
```
{
  "data": {
    "visit": {
        "id": "e36f8c3d-0052-478c-810b-521336e5aa41",
        "place_id": "2983980a-4035-4211-a98d-deb264c6f9f5",
        "visited_at": "2023-08-24T09:18:18.759584Z",
        "created_at": "2023-08-24T09:18:39.709658Z",
        "updated_at": "2023-08-24T09:18:39.709658Z"
      }
  },
  "status": "success"
}
```
## Delete A Visit By ID
Delete a visit by visit id.
**Endpoint:** `DELETE /v1/visits/visitId`
**Header** `"Authorization": "Bearer access_token"`

**Response**
```
{
	"message": "Visit successfully deleted.",
	"status": "success" 
}
```












