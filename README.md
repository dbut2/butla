# Shortener

Shortener is a simple URL shortening service written in Go. It provides the ability to create short, memorable links that redirect to longer URLs.

## Prerequisites

To run Shortener, ensure you have the following installed:
- Go (version 1.21 or newer)
- Docker (optional)

The application also requires access to a database. The provided `schema.sql` file includes the SQL schema that is compatible with MySQL.

## Building

### Manually

1. Navigate to the `shortener` directory.
2. Run the Go build command:
    ```bash
    go build -o shortener ./cmd/shortener
    ```
3. The `shortener` binary will be created in the current directory.

### Using Docker

1. Navigate to the root directory of the project.
2. Build the Docker image with the following command:
    ```bash
    docker build -t shortener -f deployment/Dockerfile .
    ```
3. The Docker image tagged as `shortener` will be created.

## Configuration

The application can be configured using the `yaml` files located in the `deployment/configs/` directory. It supports environment-specific configurations:

- `prod.yaml` for production
- `test.yaml` for testing
- `dev.yaml` for development
- `local.yaml` for local development

Adjust the configurations as needed for your environment, such as database credentials, datastore identifiers, or cache settings.

## Running

### Manually

Execute the built binary with the following command:
```bash
./shortener
```

You can set the `ENV` environment variable to load specific configurations:
```bash
ENV=local ./shortener
```

### Using Docker

1. Run the Docker image with the following command:
    ```bash
    docker run -p 8080:8080 -e ENV=local shortener
    ```
2. The application will be accessible on port `8080`.

## Deployment

Refer to the provided `Dockerfile` in `deployment/` directory for deployment configurations. Update it as necessary to use your base image and desired environment.

## Database Schema

To set up your database, use the `schema.sql` file located in the `deployment/` directory. This file contains the SQL commands to create the necessary tables and sample data.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
```