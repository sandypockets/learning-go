package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func countLinesInFile(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func getAllFilesInDir(dirname string) []string {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	var filenames []string
	for _, file := range files {
		if file.IsDir() {
			filenames = append(filenames, getAllFilesInDir(dirname+"/"+file.Name())...)
		} else {
			filenames = append(filenames, dirname+"/"+file.Name())
		}
	}
	return filenames
}

func main()  {
	startingDirPath := flag.String("dir", ".", "The directory to start counting lines of files from. Defaults to the current directory.")

	flag.Parse()

	filePaths := getAllFilesInDir(*startingDirPath)
	totalLines := 0
	for _, filePath := range filePaths {
		totalLines += countLinesInFile(filePath)
	}
	fmt.Println("TOTAL LINES: ", totalLines)

}
