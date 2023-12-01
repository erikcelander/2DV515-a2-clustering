package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Blog struct {
    Name       string
    WordCounts []int
}



func readBlogsFromFile(filename string) ([]Blog, error) {
	file, err := os.Open(filename)
	if err != nil {
			return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var blogs []Blog
	isFirstLine := true // Flag to identify the first line (header)

	for scanner.Scan() {
			if isFirstLine {
					isFirstLine = false
					continue // Skip the first line
			}

			line := scanner.Text()
			parts := strings.Split(line, "\t")

			if len(parts) < 2 { // Skip lines that don't have enough data
					continue
			}

			var wordCounts []int
			for _, countStr := range parts[1:] {
					count, err := strconv.Atoi(countStr)
					if err != nil {
							// Handle error for non-numeric values
							fmt.Printf("Error parsing word count: %s\n", countStr)
							continue
					}
					wordCounts = append(wordCounts, count)
			}

			blog := Blog{
					Name:       parts[0],
					WordCounts: wordCounts,
			}
			blogs = append(blogs, blog)
	}

	if err := scanner.Err(); err != nil {
			return nil, err
	}

	return blogs, nil
}
