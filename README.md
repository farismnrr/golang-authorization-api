# Golang Authorization API

This project provides a RESTful API for managing authorization in a Golang application.

## Features

- Authentication using JWT tokens
- CRUD operations for copyright data
- Integration with Cloudflare for protection

## Installation

1. Clone the repository:
```bash
git clone https://github.com/farismnrr/golang-authorization-api.git
```
2. Navigate to the project directory:
```bash
cd golang-authorization-api
```
3. Install dependencies:
```bash
go mod tidy
```
4. Create a `.env` file and `Authorization.json` file.
5. Run the application:
```bash
go run main.go
```
## Usage

### Endpoints

- `GET /`: Returns server information.
- `GET /api/{apiVersion}/get-key`: Retrieves the private key for authorization using middleware.
- `GET /api/{apiVersion}/copyrights`: Retrieves copyright.
- `POST /api/{apiVersion}/copyrights`: Adds new copyright.
- `PUT /api/{apiVersion}/copyrights`: Updates existing copyright.
- `DELETE /api/{apiVersion}/copyrights`: Removes copyright.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
