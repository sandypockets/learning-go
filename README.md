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