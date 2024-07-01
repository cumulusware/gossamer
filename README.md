# gossamer

Go-based website generator to create static, HTMX, and REST API sites.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE.txt]

## Overview

## Installation

```bash
$ go install github.com/cumulsware/gossamer
```

## Usage

## Documentation

Documentation can be found at either:

- <https://godoc.org/github.com/cumulusware/gossamer>
- <http://localhost:6060/pkg/github.com/cumulusware/gossamer/> after running `$
godoc -http=:6060`

## Contributing

Contributions are welcome! To contribute please:

1. Fork the repository
2. Create a feature branch
3. Code
4. Submit a [pull request][]

### Testing

Prior to submitting a [pull request][], please run:

```bash
$ just check
$ just lint
```

To update and view the test coverage report:

```bash
$ just cover
```

## License

[gossamer][] is released under the MIT license. Please see the
[LICENSE.txt][] file for more information.

[godoc badge]: https://godoc.org/github.com/cumulusware/gossamer?status.svg
[godoc link]: https://godoc.org/github.com/cumulusware/gossamer
[gossamer]: https://github.com/cumulusware/gossamer
[LICENSE.txt]: https://github.com/cumulusware/gossamer/blob/master/LICENSE.txt
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/cumulusware/gossamer
[report card]: https://goreportcard.com/report/github.com/cumulusware/gossamer
