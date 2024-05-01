# Butla Shortener

A URL shortening service.

## Overview

The Butla Shortener allows you to shorten long URLs into manageable shortened links that redirect to the original URLs. It's ideal for use in contexts where space is at a premium, such as Twitter or SMS. You can also get detailed analytics about who is clicking your links.

## Getting Started

To run the Butla Shortener service locally, follow these steps:

1. Ensure that you have Docker installed on your machine.

2. Clone this repository locally and navigate to the project's directory.

3. Build the docker container:

   ```sh
   docker build -t butla-shortener .
   ```

4. Run the docker container:

   ```sh
   docker run -p 8080:8080 butla-shortener
   ```

5. Access the service by going to [http://localhost:8080/shorten](http://localhost:8080/shorten) in your web browser.

## Features

- **Shorten URLs:** Convert long URLs into short, memorable links.
- **Redirect:** Shortened links redirect seamlessly to the targeted URLs.
- **Analytics:** Gain insights into link usage with detailed analytics.

## Configuration

Configuration is controlled through YAML files located in the `deployment/configs` directory. Different configurations are provided for production, development, testing, and local scenarios:

- `prod.yaml` - Production configuration
- `dev.yaml` - Development configuration
- `test.yaml` - Testing configuration
- `local.yaml` - Local configuration

Choose the appropriate configuration file based on your deployment needs and adjust the settings, such as the server address, database settings, cache settings, and host information.

## Contributing

Contributions to Butla Shortener are welcome. Please follow standard git workflow for contributions:

1. Fork the repository and create your feature branch.
2. Make your changes and write tests when applicable.
3. Submit a Pull Request with a comprehensive description of changes.

## License

Butla Shortener is available under the MIT license. See the LICENSE file for more info.