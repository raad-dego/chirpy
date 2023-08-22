# Chirpy Web Server

Chirpy is a web server application built using Go that provides an API for creating, retrieving, and managing chirps.

## Getting Started

### Prerequisites

- Go (1.13 or higher)
- Git

### Installation

1. Clone the repository:
git clone https://github.com/raad-dego/chirpy.git

2. Navigate to the project directory:
cd chirpy

3. Create a `.env` file in the project directory and set the required environment variables:

4. Build and run the application:
go build -o out && ./out

## Usage

The Chirpy web server provides endpoints to manage chirps and users.

### API Endpoints

- **GET /api/chirps**: Retrieve chirps.
- **GET /api/chirps/{chirpID}**: Retrieve a specific chirp by ID.
- **POST /api/chirps**: Create a new chirp.
- **DELETE /api/chirps/{chirpID}**: Delete a chirp.

- **POST /api/users**: Create a new user.
- **PUT /api/users**: Update a user's information.

- **POST /api/login**: Authenticate user login and generate JWT.
- **POST /api/revoke**: Revoke a JWT.
- **POST /api/refresh**: Refresh an expired JWT.

- **POST /api/polka/webhooks**: Handle webhook for Polka verification.

- **GET /admin/metrics**: Retrieve server metrics.

## Configuration

The application can be configured using environment variables in the `.env` file.

- `JWT_SECRET`: Secret key for JWT generation and validation.
- `POLKA_API`: URL for the Polka API.

## Contributing

Feel free to contribute to this project by submitting pull requests or issues.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
