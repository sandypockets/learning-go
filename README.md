# Learning Go
A small collection of simple programs for getting familiar with Golang

## CSV Quiz
A tiny CLI that reads problems and answers from a CSV file, and prints them out in a quiz format for the user to solve on the command line.

### Packages
- [encoding/csv](https://golang.org/pkg/encoding/csv/)
- [flag](https://golang.org/pkg/flag/)
- [fmt](https://golang.org/pkg/fmt/)
- [os](https://golang.org/pkg/os/)
- [strings](https://golang.org/pkg/strings/)

### Usage
1. Clone the repository.
2. Change into the `csvQuiz` directory.
3. Build the binary executable: `go build csvQuiz.go`
4. Run the program: `./csvQuiz` (use the `-csv=somefilename.csv` flag to specify a different CSV file)
5. If you need help, run the program with the `-h` or `--help` flag.

## URL Shortener
A super simple server that checks if a URL exists for a given key, and if it does, redirects to the URL. This program is extremely basic; it doesn't actually shorten URLs, and is more of an implementation of the handler. 

### Packages
- [fmt](https://golang.org/pkg/fmt/)
- [net/http](https://golang.org/pkg/net/http/)

### Usage
1. Clone the repository.
2. Change into the `urlShortener` directory.
3. Run `go run urlShortener.go` to start the server.
4. Visit `http://localhost:8080` in your browser.
5. Visit `http://localhost:8080/github` in your browser to see the redirect.

## Count Lines
A simple program that counts the number of lines in a file, or project. If the starting directory contains other directories, the program will run recursively until all file lines are counted. 

### Packages
- [bufio](https://golang.org/pkg/bufio/)
- [flag](https://golang.org/pkg/flag/)
- [fmt](https://golang.org/pkg/fmt/)
- [io/ioutil](https://golang.org/pkg/io/ioutil/)
- [os](https://golang.org/pkg/os/)

### Usage
1. Clone the repository.
2. Change into the `countLines` directory.
3. Build the executable: `go build countLines.go`
4. Run the executable: `./countLines --dir=some/directory`. If you do not specify a directory, the program will default to the current directory.
