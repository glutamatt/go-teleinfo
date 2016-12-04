package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func readFile(path string, buf chan string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)

	for {
		if scanner.Scan() {
			if text := scanner.Text(); text != "" {
				buf <- text
			}
		} else {
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func parseLines(lines chan string, mesures chan int, indices chan int) {
	for line := range lines {
		if chunks := strings.Split(line, " "); len(chunks) == 3 {
			if chunks[0] == "PAPP" {
				if val, err := strconv.ParseInt(chunks[1], 10, 32); err == nil {
					mesures <- int(val)
				}
			}
			if chunks[0] == "BASE" {
				if val, err := strconv.ParseInt(chunks[1], 10, 32); err == nil {
					indices <- int(val)
				}
			}
		}
	}
}
