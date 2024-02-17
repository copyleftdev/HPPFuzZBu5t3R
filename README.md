# HPPFuzZBu5t3R

HPPFuzZBu5t3R is a command-line tool designed to detect HTTP Parameter Pollution (HPP) vulnerabilities in web applications. By exploiting the way web applications parse query parameters, HPPFuzZBu5t3R helps security researchers identify potential security flaws that could be used to manipulate web application logic or bypass security controls.

## Features

- Supports testing for both comma-separated values and file-based input for parameters.
- Utilizes concurrent requests for efficient scanning.
- Color-coded output for easy interpretation of results.
- Customizable query parameter and target URL inputs.

## Installation

Before you begin, ensure you have [Go](https://golang.org/dl/) installed on your system.

To install HPPFuzZBu5t3R, clone the repository and build the tool with Go:

```bash
git clone https://github.com/copyleftdev/HPPFuzZBu5t3R.git
cd HPPFuzZBu5t3R
go build -o hppfuzzbu5t3r
```

## Usage

Run HPPFuzZBu5t3R with the following command-line arguments:

- `--target` (or `-t`): Specifies the target URL to test.
- `--param` (or `-p`): Specifies the query parameter to test for HPP.
- `--data` (or `-d`): Specifies the values to test, either as a comma-separated list or a file path to a newline-separated list.

### Example

Testing with comma-separated values:

```bash
./hppfuzzbu5t3r --target "http://example.com" --param "search" --data "safeValue,' OR 1=1;--"
```

Testing with file input:

```bash
./hppfuzzbu5t3r --target "http://example.com" --param "search" --data "./values.txt"
```

