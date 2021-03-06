package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/milanaleksic/logtime"
	"log"
	"os"
)

var (
	logTimeLayout string
	inputFile     *os.File
)

func init() {
	var logTime string
	var inputFileLocation string
	flag.StringVar(&logTime, "log-time", `2006-01-02 15:04:05`, "pattern that should match beginning of all log lines")
	flag.StringVar(&inputFileLocation, "input-file", "", "which file to process (default - stdin)")
	flag.Parse()

	logTimeLayout = logTime
	if inputFileLocation != "" {
		var err error
		inputFile, err = os.Open(inputFileLocation)
		if err != nil {
			log.Fatalf("Failed to open input file: %s, reason: %v", inputFileLocation, err)
		}
	}
}

func main() {
	var scanner *bufio.Scanner
	if inputFile != nil {
		scanner = bufio.NewScanner(inputFile)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
		_, _ = fmt.Fprintln(os.Stderr, "Reading from stdin")
	}
	moments := logtime.NewLogTime(logTimeLayout).ReadStreamOfLogLines(scanner)
	for _, moment := range *moments {
		fmt.Printf("%f\t%s\n", moment.Duration.Seconds(), moment.Line)
	}
}
