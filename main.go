package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func loadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "=")
		if len(splitLine) != 2 {
			return fmt.Errorf("invalid env line: %s, expecting x=y", line)
		}
		err := os.Setenv(splitLine[0], splitLine[1])
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func main() {
	args := os.Args[1:] // Starting from 1 ignoring the binary name
	filePath := args[0]
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("darwin")
		fmt.Println("Not supported yet")
	case "linux":
		fmt.Println("linux")
		err := loadEnv(filePath)
		if err != nil {
			fmt.Println("Error loading env:", err)
		}
	case "windows":
		fmt.Println("windows")
		err := loadEnv(filePath)
		if err != nil {
			fmt.Println("Error loading env:", err)
		}
	default:
		fmt.Println("IDK what you are using buddy")
	}
}
