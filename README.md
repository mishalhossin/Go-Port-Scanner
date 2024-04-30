# Go Port Scanner

This is a simple port scanner written in Go, capable of scanning TCP ports on a target IP address.

## How It Works

The port scanner utilizes Go's concurrency features to scan a range of TCP ports on a target IP address simultaneously. It iterates through ports 1 to 65535, attempting to establish a TCP connection with each port within a specified timeout duration. If a connection is successful, the scanner reports that the port is open.

## Features

- Concurrent scanning: Utilizes goroutines for concurrent port scanning, improving efficiency.
- Timeout handling: Allows specifying a timeout duration for connection attempts to prevent long delays.
- Comprehensive scanning: Scans ports 1 to 65535 to provide a comprehensive overview of open ports on the target IP address.
- User-friendly interface: Provides prompts for input and clear output for easy interaction.

## Contributing

Contributions are welcome! Feel free to open issues for bug reports, feature requests, or submit pull requests for enhancements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
