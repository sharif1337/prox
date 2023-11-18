package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Define flags
	appendFlag := flag.Bool("a", false, "Append the replacement string")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: go run prox.go [-a] <replacement>")
		return
	}

	replacement := args[0]
	fileContent, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	modifiedContent := modifyURLs(string(fileContent), replacement, *appendFlag)
	fmt.Println(modifiedContent)
}

func modifyURLs(input, replacement string, appendFlag bool) string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		lines[i] = modifyURL(line, replacement, appendFlag)
	}

	return strings.Join(lines, "\n")
}

func modifyURL(urlString, replacement string, appendFlag bool) string {
	parts := strings.Split(urlString, "?")
	if len(parts) != 2 {
		return urlString
	}

	baseURL, queryString := parts[0], parts[1]

	// Update parameters without encoding
	params := strings.Split(queryString, "&")
	for i := range params {
		pair := strings.SplitN(params[i], "=", 2)
		if len(pair) == 2 {
			if appendFlag {
				pair[1] += replacement
			} else {
				pair[1] = replacement
			}
			params[i] = strings.Join(pair, "=")
		}
	}

	// Reconstruct the modified URL
	return baseURL + "?" + strings.Join(params, "&")
}
