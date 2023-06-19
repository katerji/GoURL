# GoURL API

This repository contains the code for the GoURL API, which is built using the Gin framework. The API provides endpoints for URL shortening and user authentication.

## API Endpoints

### Router Handler

- Endpoint: `GET /r/*param`

   This endpoint handles URL redirection. It extracts the shortened URL from the request path and redirects the user to the associated destination URL.

### URL Delete Handler

- Endpoint: `DELETE /url`

   This endpoint deletes a URL record. It expects a JSON payload with an `id` field specifying the URL ID to be deleted.

### URL Create Handler

- Endpoint: `POST /url`

   This endpoint creates a new shortened URL. It expects a JSON payload with an `original_url` field specifying the original URL to be shortened. The user must be authenticated to access this endpoint.

For detailed information about the request payloads and response formats, please refer to the code comments and implementations in the respective handler files.

## Authentication

The GoURL API includes authentication functionality. To access the authenticated endpoints, you need to include an authentication token in the request headers. The token is obtained by logging in or registering a new user.

## Contact

For any questions or inquiries, you can reach out to the project maintainer:

- GitHub: [katerji](https://github.com/katerji)

Please note that while the maintainer will try to address any issues or inquiries, there may be no guarantee of a timely response.
