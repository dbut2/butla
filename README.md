# Shortener Service

This service provides functionality to shorten URLs and manage shortened URL redirection. 

## Getting Started

### Prerequisites

- Go 1.21
- MySQL (for persistent storage)
- Redis (for caching layer)
- Docker (for containerization and deployment)

### Configuration

The service can be configured using a YAML configuration file. Example files for different environments are provided in the `deployment/configs/` directory. Customize the `local.yaml` file for local development or create a new configuration for your environment.

### Running Locally

For local development, ensure that the MySQL and Redis instances are accessible and the `local.yaml` configuration file has the correct connection details.

### Building the Service

You can use the provided `Dockerfile` in the `deployment` directory to build the service:

```sh
cd deployment
docker build -t shortener .
```

### Running the Service

To run the built Docker image:

```sh
docker run -p 8080:8080 shortener
```

The API will be available at `http://localhost:8080`.

## API Endpoints

- `GET /`: Redirects to the default shortened URL.
- `GET /:code`: Given a shortened code, it redirects to the original URL if valid and not expired.
- `POST /shorten`: Shortens a given URL.

## Database Schema

The MySQL database schema can be found in `deployment/database/schema.sql`. This must be applied to the MySQL instance before starting the service.

## Deployment

Modify and use the Dockerfile and Kubernetes configuration files from the `deployment` directory for deploying the service to your preferred environment.

## Contributing

Please submit any bugs or feature requests as issues on the repository.

For contributing code, fork the repository, make your changes, and open a pull request.

## License

Specify the license under which the project is made available.
```