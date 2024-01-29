# HTTPedia [![build-release-binary](https://github.com/rnemeth90/HTTPedia/actions/workflows/build.yaml/badge.svg)](https://github.com/rnemeth90/HTTPedia/actions/workflows/build.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/rnemeth90/HTTPedia)](https://goreportcard.com/report/github.com/rnemeth90/HTTPedia)

## Overview
HTTPedia is a command-line tool designed to provide descriptions and details for HTTP status codes. It's an invaluable resource for developers and system administrators for quick reference and debugging HTTP responses.

## Features
- Retrieve descriptions for HTTP status codes.
- Simple and intuitive command-line interface.
- Extensible to include additional details about each status code in future updates.

## Getting Started

## Dependencies
**To build from source, Go v1.13+ is required.**

## Installing
Download the latest release [here](https://github.com/rnemeth90/HTTPedia/releases)

## Usage
To use HTTPedia and retrieve information about a specific HTTP status code:
```
$ httpedia -s 404
```
HTTPedia will provide a description of the HTTP 404 status code.

## Example Output
```
$ httpedia -s 404
Not Found
```

## Contributing
Contributions to HTTPedia are welcome! Please refer to our [contribution guidelines](CONTRIBUTING.md) for more details on how to report bugs, suggest features, or submit pull requests.

## To Do
- [ ] Implement basic functionality for HTTP status code lookup.
- [ ] Extend the database to include less common status codes.
- [ ] Enhance output formatting and add additional information for each status code.
- [ ] Add support for custom status code databases.

## Version History
* 1.0.0
    * Initial Release

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details
