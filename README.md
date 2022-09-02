# termcolor

![CI Status][ci-img-url] [![Go Report Card][go-report-card-img-url]][go-report-card-url]
[![Package Doc][package-doc-img-url]][package-doc-url] [![Releases][release-img-url]][release-url]

[ci-img-url]: https://github.com/halimath/termcolor/workflows/CI/badge.svg
[go-report-card-img-url]: https://goreportcard.com/badge/github.com/halimath/termcolor
[go-report-card-url]: https://goreportcard.com/report/github.com/halimath/termcolor
[package-doc-img-url]: https://img.shields.io/badge/GoDoc-Reference-blue.svg
[package-doc-url]: https://pkg.go.dev/github.com/halimath/termcolor
[release-img-url]: https://img.shields.io/github/v/release/halimath/termcolor.svg
[release-url]: https://github.com/halimath/termcolor/releases

`termcolor` provides a simple and convenient API to output colorized terminal output with support for
terminal detection and color supression.

# Usage

`termcolor` provides a `Printer` type which accepts strings and optional style attributes and outputs 
colorized output to an `io.Writer`. Factory functions for `STDOUT`, `STDERR` and `*os.File` are included which
automatically detect, if the underlying device is a character device.

```go
p := termcolor.Stdout()

p.Printf("Welcome to %s output!", p.Styled("colored", termcolor.ForegroundCyan), termcolor.Bold)
```

## Installation

`termcolor` uses go modules and requires Go 1.11 or greater.

```
$ go get -u github.com/halimath/termcolor
```

# Changelog

## 0.1.0
* Initial release

# License

```
Copyright 2022 Alexander Metzner.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
