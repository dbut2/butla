# Shortener Service

## Overview

The Shortener Service is a URL shortening service with the capability to create shortened URLs (links) and redirect from a short link to the original URL. It also supports user registration and authentication for managing links.

## Service Components

- **Shortener**: Provides functionalities to shorten long URLs, manage and retrieve the shortened links, handles the short link redirection to the original URL, and ensures the links may expire after a certain time or restricted based on the requester's IP.
- **Auth**: Handles user registration and authentication services.
- **Server**: Powers the HTTP server that exposes endpoints for the services provided by the application and serves the related web pages for user interaction.

## Configuration

The service can be configured using YAML configuration files. Configuration files for different environments (local, development, production, and test) can be found in the `deployment/configs/` directory:

- `local.yaml` - Configuration for local development setup.
- `dev.yaml` - Configuration for the development environment.
- `prod.yaml` - Configuration for the production environment.
- `test.yaml` - Configuration for the test environment.

The service can be customized by editing the respective environment's configuration file or setting the `ENV` or `CONFIG_FILE` environment variables to point to a custom configuration file.

### Configuration Parameters

- `address`: The address the server listens on.
- `store`: Specifies the type and configuration for the storage mechanism used for persisting links and users.
  - `database`: Parameters for the database connectivity.
  - `datastore`: Parameters for Google Cloud Datastore configuration.
  - `cache`: Configuration for caching layer (using Redis).
- `host`: Details about the hostname and scheme.

## Database Schema

Database `shortener`, which contains two main tables:

- **links**: Stores the short links created by the service.
- **users**: Stores the user accounts for the authorization and management of links.

The schema definition can be found in `deployment/database/schema.sql`. It includes the table creation statements and initial data setup.

## Deployment

The service is containerized using Docker. The `Dockerfile` in the `deployment/` directory can be used to build the service image. The image build process contains the following stages:

1. Building the Go application from source.
2. Copying the built application into a clean image for deployment.

Set the `ENV` variable during the build or run phase to customize the environment settings for the container.

## Running the Service

Execute the following command to start the service using Docker:

```sh
docker build -t shortener-service . -f deployment/Dockerfile
docker run -p 8080:8080 -e ENV=prod shortener-service
```

Replace `prod` with your desired environment (`local`, `dev`, `test`, or path to your custom configuration YAML).

Once the service is running, access it via the configured hostname, or `localhost:8080` if running locally.

## API Endpoints

- `GET /shorten`: Serves the page to create a shortened URL.
- `POST /shorten`: API endpoint to create a shortened URL.
- `GET /`: Redirects to the default link if set.
- `GET /{code}`: Redirects to the original URL corresponding to the given `code`.

## User Registration and Authentication

User registration and authentication functionalities are available through the Auth interface but are not exposed via the HTTP server in the current implementation. To utilize these functionalities, additional endpoints should be implemented in the server component.

---

This README is part of the technical documentation for the Shortener Service.