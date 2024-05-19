package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// >ccwc -c test.txt
//   342190 test.txt

func main() {
	// Parse the command line arguments
	countBytes := flag.Bool("c", false, "Count the number of bytes in file")
	countWords := flag.Bool("w", false, "Count the number of words in file")
	countLines := flag.Bool("l", false, "Count the number of lines in file")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	fileName := args[0]

	// Read the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	var count int

	// Count the number of bytes/words/lines in file
	if *countBytes {
		count, err = CountFileByBytes(file)
		if err != nil {
			fmt.Printf("Error counting bytes: %v\n", err)
			os.Exit(1)
		}

	} else if *countWords {
		count, err = CountFileByWords(file)
		if err != nil {
			fmt.Printf("Error counting words: %v\n", err)
			os.Exit(1)
		}
	
	} else if *countLines{
		count, err = CountFileByLines(file)
		if err != nil {
			fmt.Printf("Error counting lines: %v\n", err)
			os.Exit(1)
		}
	} else{
		byteCount, err := CountFileByBytes(file)
		if err != nil {
			fmt.Printf("error counting files: %v\n", err)
			os.Exit(1)
		}

		_, err = file.Seek(0,0)
		if err != nil {
			fmt.Printf("Error resetting file: %v\n", err)
		}
		wordCount, err := CountFileByWords(file)
		if err != nil {
			fmt.Printf("error counting words: %v\n", err)
			os.Exit(1)
		}

		_, err = file.Seek(0,0)
		if err != nil {
			fmt.Printf("Error resetting file: %v\n", err)
		}
		lineCount, err := CountFileByLines(file)
		if err != nil {
			fmt.Printf("error couting lines: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("  %d  %d  %d  %s\n", lineCount, wordCount, byteCount, fileName)
		return 
	}
	fmt.Printf("%d %s\n", count, fileName)
}

func CountFileByBytes(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)

	count := 0
	for scanner.Scan() {
		count++
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}
	return count, nil
}

func CountFileByWords(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}
	return count, nil
}

func CountFileByLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}
	return count, nil
}
