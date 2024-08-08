# CONTACT-MAIL-GO

This backend service is designed to handle various tasks related to email communication and user authentication. It provides features such as email validation, HTML email template parsing, and user authentication with 2-phase verification.

## Features

- **Email Validation**: Verifies email addresses for syntax correctness, domain existence, and whether the email is disposable or a role account.
- **HTML Template Parsing**: Renders and sends HTML emails using customizable templates.
- **2-Phase Authentication**: Secure authentication mechanism involving token generation and validation for API access.

## UPCOMING IMPORVEMENTS

- \*\*Added support for multiple templates and user templates.
- \*\*Increased support for multiple SMTP Servers.
- \*\*Web Interface to get your personal singing secret to use the webAPI.

## Table of Contents

1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Configuration](#configuration)
5. [API Endpoints](#api-endpoints)
6. [Usage as your own service](#usage)
7. [License](#license)

## Overview

The backend service is built with Go and uses the Chi router for handling HTTP requests. It integrates with an email verification library and supports sending emails with dynamic HTML templates.

## Prerequisites

Before running the backend service, ensure you have:

- [Go](https://golang.org/dl/) installed on your machine.
- Access to an SMTP server for sending emails.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/backend-service.git
   ```

2. Navigate to the backend directory:

   ```bash
   cd backend-service
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Configuration

Create a `.env` file in the root of the backend directory with the following content:

```env
SECRET=your-secret-key
JWTKEY=your-jwt-secret-key
```

Replace your-secret-key with your secret. It should be a secure random string used for token generation.

## API Endpoints

1. Token Generation
   Endpoint: GET /api/getToken

   Description: Generates and returns a token for API authentication.

   BODY:
   {
   "secret":"conoienoqeiiqeno",
   "smtp_user":"xyzgamer@gmail.com",
   "smtp_pass":"asfc poqw kdnv adfa"
   }
   Responses:

   200 OK: {
   "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzwwencowenciewncoinopcmVkQXQiOiIxNjg0Mzg5wovineoinvweoinvowenvoeionjAwfQ.S4ydJ...QhTk"
   } If the token is successfully generated.
   401 Unauthorized: If the secret key is incorrect.

2. Send Email
   Endpoint: POST /api/sendEmail

   Description: Receives email data, validates the email address, and sends a personalized HTML email.

   Request Body:

   {
   "email":"xyzgamer@gmail.com",
   "message":"Just testing this",
   "receiver_name":"Aviral",
   "sender_name":"Aviral Shukla",
   "subject":"Hey! This is a testing message"
   }

   Headers:
   Authorization: Bearer <token>

   Responses:
   201 Created: If the email is successfully sent.
   400 Bad Request: If the email address is invalid or required fields are missing.
   401 Unauthorized: If the token is invalid.
   500 Internal Server Error: For any server-side errors.

## Usage as your own service

    1. Start the Backend Server:

    go run main.go

    2. Get a Token:

    Make a GET request to /api/token with the correct Authorization header to obtain an API token.

    3. Send an Email:

    Make a POST request to /api/sendEmail with the required JSON body and the token in the Authorization header.

    Example Request to Get Token:

    ```bash
    curl -X POST http://localhost:8080/api/getToken \
    -H "Content-Type: application/json" \
    -d '{
         "secret":"conoienoqeiiqeno",
         "smtp_user":"xyzgamer@gmail.com",
         "smtp_pass":"asfc poqw kdnv adfa"
    }'
    ```

    Example Request to Send Email

    ```bash
    curl -X POST http://localhost:8080/api/sendEmail \
    -H "Authorization: Bearer your-token" \
    -H "Content-Type: application/json" \
    -d '{
         "email":"xyzgamer@gmail.com",
         "Message":"Just testing this",
         "receiver_name":"Aviral",
         "sender_name":"Aviral Shukla",
         "subject":"Hey! This is a testing message"
    }'
    ```

# User will receive a mail like below:

<img width="1260" alt="image" src="https://github.com/user-attachments/assets/7b288d8c-9b8d-4faa-ae56-1daa328d61f7">

# Feel free to contribute
