# URL Shortener Service

This project is a simple URL shortening service built with Go. It provides a straightforward REST API for shortening URLs and redirecting shortened URLs back to their original form.

## Features

- Shorten long URLs via POST request
- Redirect shortened URLs to original URLs
- Simple error handling with proper HTTP status codes

## Endpoints

### Shorten URL

- **Method:** `POST`
- **Path:** `/`
- **Request Body:** Original URL as plain text (`text/plain`)
- **Response:** Shortened URL as plain text (`text/plain`)

#### Example Request:
```http
POST / HTTP/1.1
Host: localhost:8080
Content-Type: text/plain

https://practicum.yandex.ru/
```

#### Example Response:
```http
HTTP/1.1 201 Created
Content-Type: text/plain
Content-Length: 30

http://localhost:8080/EwHXdJfB
```

### Redirect URL

- **Method:** `GET`
- **Path:** `/{id}` (e.g., `/EwHXdJfB`)
- **Response:** HTTP redirect (307) to the original URL

#### Example Request:
```http
GET /EwHXdJfB HTTP/1.1
Host: localhost:8080
```

#### Example Response:
```http
HTTP/1.1 307 Temporary Redirect
Location: https://practicum.yandex.ru/
```

### Error Handling

On invalid requests, the service responds with:
```http
HTTP/1.1 400 Bad Request
```

## Getting Started

### Requirements

- Go 1.24.0

### Installation

Clone the repository:

```bash
git clone https://github.com/Kudzeri/yandex-practicum-url-shortener.git
cd yandex-practicum-url-shortener
```

### Running the Server

```bash
go run .
```

The server will be accessible at [http://localhost:8080](http://localhost:8080).


## Author

- GitHub: [Kudzeri](https://github.com/Kudzeri)

## License

This project is open-source and available under the MIT License.