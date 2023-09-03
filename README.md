# task-5-pbi-btpns-RoniRagilImanKhoirul

This repository showcases my final project for the Fullstack Developer Virtual Internship Experience at BTPN Syariah. As a participant, I have developed a comprehensive web application that serves as my demonstration of skills and knowledge in fullstack development. In this project, I have built a robust backend application using the Go programming language (Golang). This challenging task has allowed me to design, develop, and deploy a fully functional API, equipping me with valuable skills for future endeavors in the field of backend development

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Authentication](#authentication)

## Installation

1. Clone the repository:

```shell
git clone https://github.com/roniragilimankhoirul/task-5-pbi-btpns-RoniRagilImanKhoirul.git
```

2. Navigate to the Project Directory:

```shell
cd task-5-pbi-btpns-RoniRagilImanKhoirul
```

3. Set Up Environment Variables:

Create a .env file in the project root and configure your environment variables.

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=rakamin
```

4. Install dependencies:

```shell
go mod tidy
```

3. Start the server:

```shell
go run main.go
```

The application should now be running on http://localhost:8080.

## Usage

### Home Endpoint

- **Method:** GET
- **URL:** `/`
- **Description:** Get a welcome message.
- **Response:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    {
      "message": "Halo, Rakamin!"
    }
    ```

### User Registration Endpoint

- **Method:** POST
- **URL:** `/users/register`
- **Description:** Register a new user.
- **Request Body:**
  ```json
  {
    "username": "test69",
    "email": "test69@test.com",
    "password": "testtest"
  }
  ```
- **Respond:**

  - **Status Code:** 200 OK
  - **Body:**

    ```json
    {
      "message": "Successfully registered",
      "data": {
        "email": "test69@test.com",
        "id": "4cc53717-100d-48b6-aaf2-97bbb4e9a8c9",
        "username": "test69"
      }
    }
    ```

### User Login Endpoint

- **Method:** POST
- **URL:** `/users/login`
- **Description:** Login a user.
- **Request Body:**
  ```json
  {
    "username": "test69",
    "email": "test69@test.com",
    "password": "testtest"
  }
  ```
- **Respond:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    {
      "message": "Successfully login",
      "data": {
        "email": "test69@test.com",
        "id": "4cc53717-100d-48b6-aaf2-97bbb4e9a8c9",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM4MzAwMzgsInN1YiI6IjRjYzUzNzE3LTEwMGQtNDhiNi1hYWYyLTk3YmJiNGU5YThjOSJ9.nzG3u3wbtQMvtQJZZTHNifbJXRjlmSKCq0qHaZV6Ldk",
        "username": "test69"
      }
    }
    ```
- **Cookie:** Authorization=token (set with a 1-day expiration)

### User Update Endpoint

- **Method:** PUT
- **URL:** `/users/{userId}`
- **Description:** Update a user.
- **Request Body:**
  ```json
  {
    "username": "test69",
    "email": "test69@test.com",
    "password": "testtest"
  }
  ```
- **Respond:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    {
      "message": "Successfully updated",
      "data": {
        "email": "test69@test.com",
        "id": "4cc53717-100d-48b6-aaf2-97bbb4e9a8c9",
        "username": "test69"
      }
    }
    ```

### User Delete Endpoint

- **Method:** DELETE
- **URL:** `/users/{userId}`
- **Description:** Delete a user.
- **Response:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    {
      "message": "Successfully deleted"
    }
    ```

### Create a New Photo

- **Endpoint:** `/api/photos`
- **Method:** `POST`
- **Description:** Create a new photo.
- **Request Body:**
  ```json
  {
    "title": "Sample Photo",
    "caption": "This is a test photo.",
    "photourl": "https://example.com/sample.jpg"
  }
  ```
- **Respond:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    {
      "data": {
        "id": "f0fd7d0e-6b00-4ae2-a2bd-796eb05794d8",
        "title": "Sample Photo",
        "caption": "This is a test photo.",
        "photourl": "https://example.com/sample.jpg",
        "userid": "4cc53717-100d-48b6-aaf2-97bbb4e9a8c9"
      },
      "message": "Photo has been created successfully"
    }
    ```

### Get User's Photos

- **Endpoint:** `/api/photos`
- **Method:** `GET`
- **Description:** Get a list of photos owned by the authenticated user.
- **Response:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    {
      "data": [
        {
          "id": "3500a952-a0a5-4bd0-b157-ae4e14b37a35",
          "title": "Sample Photo",
          "caption": "This is a test photo.",
          "photourl": "https://example.com/sample.jpg",
          "userid": "4cc53717-100d-48b6-aaf2-97bbb4e9a8c9"
        },
        {
          "id": "f0fd7d0e-6b00-4ae2-a2bd-796eb05794d8",
          "title": "Sample Photo",
          "caption": "This is a test photo.",
          "photourl": "https://example.com/sample.jpg",
          "userid": "4cc53717-100d-48b6-aaf2-97bbb4e9a8c9"
        }
      ],
      "message": "success"
    }
    ```

### Update a Photo

- **Endpoint:** `/api/photos/{photoId}`
- **Method:** `PUT`
- **Description:** Update an existing photo.
- **Request Body:**
  ```json
  {
    "title": "Sample Photo XXX",
    "caption": "This is a test photo XXX",
    "photourl": "https://example.com/sample-XXX.jpg"
  }
  ```
- **Respond:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    {
      "data": {
        "id": "3500a952-a0a5-4bd0-b157-ae4e14b37a35",
        "title": "Sample Photo XXX",
        "caption": "This is a test photo XXX",
        "photourl": "https://example.com/sample-XXX.jpg",
        "userid": "4cc53717-100d-48b6-aaf2-97bbb4e9a8c9"
      },
      "message": "data has been updated"
    }
    ```

### Delete a Photo

- **Endpoint:** `/api/photos/{photoId}`
- **Method:** `DELETE`
- **Description:** Delete an existing photo.
- **Response:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    {
      "message": "Deleted successfully"
    }
    ```

## Authentication

- User registration and login endpoints do not require authentication.

- All other endpoints require authentication via the `Authorization` header with the JWT token obtained during login.
