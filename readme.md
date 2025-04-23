# Bookstore Golang

This is a simple bookstore application built with Golang and Gin framework. It includes features for user authentication, book management, and more.

## Installation

Follow these steps to set up the application:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/emkr-13/bookstore_go_v1.git
   cd bookstore_go_v1
   ```

2. **Install dependencies**:
   Ensure you have Go installed (version 1.18 or later). Then, run:

   ```bash
   go mod tidy
   ```

3. **Set up the database**:

   - Install and configure a database (e.g., PostgreSQL, MySQL, or SQLite).
   - Update the database configuration in `config/config.go` to match your database credentials.

4. **Run database migrations**:
   The application will automatically migrate the database models when you start the server.

5. **Run the application**:

   ```bash
   go run cmd/main.go
   ```

6. **Access the application**:
   The server will start on the port specified in the configuration (default: `8080`). Open your browser or use a tool like Postman to access:
   ```
   http://localhost:8080/api/v1
   ```

## Usage

### Public Endpoints

These endpoints do not require authentication:

- **Register**: `POST /api/v1/register`
  - Request body:
    ```json
    {
      "username": "your_username",
      "password": "your_password"
    }
    ```
- **Login**: `POST /api/v1/login`
  - Request body:
    ```json
    {
      "username": "your_username",
      "password": "your_password"
    }
    ```
- **Refresh Token**: `POST /api/v1/refresh-token`
  - Request body:
    ```json
    {
      "refresh_token": "your_refresh_token"
    }
    ```

### Protected Endpoints

These endpoints require authentication (JWT token in the `Authorization` header):

- **Logout**: `POST /api/v1/logout`
- **Create Book**: `POST /api/v1/books`
  - Request body:
    ```json
    {
      "title": "Book Title",
      "author_id": 1,
      "publisher_id": 1,
      "price": 19.99
    }
    ```
- **Get All Books**: `GET /api/v1/books`
- **Get Book by ID**: `GET /api/v1/books/:id`
- **Update Book**: `PUT /api/v1/books/:id`
  - Request body:
    ```json
    {
      "title": "Updated Title",
      "author_id": 1,
      "publisher_id": 1,
      "price": 24.99
    }
    ```
- **Delete Book**: `DELETE /api/v1/books/:id`

## Configuration

The application uses a configuration file located in `config/config.go`. Update the following fields as needed:

- `DB` (database connection)
- `AppPort` (server port)
- `JWTSecret` (secret key for JWT)
- `AuthExp` (authentication token expiration time)
- `RefreshExp` (refresh token expiration time)

## Contributing

Feel free to fork this repository and submit pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License.
