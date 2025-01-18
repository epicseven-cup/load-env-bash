package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var inputFileName string
var outFileName string

func init() {
	const (
		defaultInputFilePath  = "./.env"
		defaultOutputFilePath = "./output.sh"
	)
	flag.StringVar(&inputFileName, "i", defaultInputFilePath, "Path to the input file (shorthand)")
	flag.StringVar(&inputFileName, "input", defaultInputFilePath, "Path to the input file")

	flag.StringVar(&outFileName, "o", defaultOutputFilePath, "Path to the output file (shorthand)")
	flag.StringVar(&outFileName, "output", defaultOutputFilePath, "Path to the output file")
	flag.Parse()
}

func loadEnv() error {
	file, err := os.OpenFile(inputFileName, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	outputFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(outputFile)
	defer file.Close()
	defer outputFile.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "=")
		if len(splitLine) != 2 {
			return fmt.Errorf("invalid env line: %s, expecting x=y", line)
		}
		_, err := writer.WriteString(fmt.Sprintf("export %s=%s\n", splitLine[0], splitLine[1]))
		if err != nil {
			return err
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}
	fmt.Println("Load env succeeded")
	return nil
}

func main() {
	switch goos := runtime.GOOS; goos {
	case "darwin":
		fmt.Println("OS: darwin")
		fmt.Println("Not supported yet")
	case "linux":
		fmt.Println("OS: linux")
		err := loadEnv()
		if err != nil {
			fmt.Println("Error loading env:", err)
		}
	case "windows":
		fmt.Println("OS: windows")
		err := loadEnv()
		if err != nil {
			fmt.Println("Error loading env:", err)
		}
	default:
		fmt.Println("IDK what you are using buddy")
	}
}
