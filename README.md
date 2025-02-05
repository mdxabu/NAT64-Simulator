# NAT64 Simulator

## Overview

The NAT64 Simulator is a tool designed to simulate NAT64 (Network Address Translation from IPv6 to IPv4) environments. This can be useful for testing and development purposes, allowing developers to ensure their applications work correctly in NAT64 scenarios.

## Features

- Simulate NAT64 translation
- Test IPv6-only environments
- Monitor and log network traffic
- Configurable settings for various scenarios

## Installation

To install the NAT64 Simulator, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/mdxabu/nat64-simulator.git
    ```
2. Navigate to the project directory:
    ```sh
    cd nat64-simulator
    ```
3. Install the required dependencies:
    ```sh
    go mod tidy
    ```

## Usage

To start the NAT64 Simulator, run the following command:
```sh
go run main.go
```


## Contributing

We welcome contributions! Please read our [contributing guidelines](CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Contact

For any questions or feedback, please open an issue on the [GitHub repository](https://github.com/mdxabu/nat64-simulator/issues).
