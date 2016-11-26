package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var htmlTemplate = "<!DOCTYPE html><html><head><link rel=\"stylesheet\"" +
	"href=\"//cdn.jsdelivr.net/chartist.js/latest/chartist.min.css\">" +
	"<script src=\"//cdn.jsdelivr.net/chartist.js/latest/chartist.min.js\"></script>" +
	"</head><body><div class=\"ct-chart ct-minor-sixth\"></div><script>setTimeout(function(){location = ''},5000);" +
	"var data = {series: [[%s]]}; new Chartist.Line('.ct-chart', data, {low: 0,showArea: true, showLine: false, showPoint: false});" +
	"</script></body></html>"

func main() {
	valuesCount := 100
	values := []int{}

	fileStrings := make(chan string)
	measures := make(chan int)
	go readFile("/dev/ttyS0", fileStrings)
	go parseLines(fileStrings, measures)
	go runHTTPServer(&values)

	for m := range measures {
		values = append(values, m)
		if len(values) > valuesCount {
			values = values[1:valuesCount]
		}
	}
}

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

func parseLines(lines chan string, mesures chan int) {
	for line := range lines {
		if chunks := strings.Split(line, " "); len(chunks) == 3 {
			if chunks[0] == "PAPP" {
				if val, err := strconv.ParseInt(chunks[1], 10, 32); err == nil {
					mesures <- int(val)
				}
			}
		}
	}
}

func runHTTPServer(values *[]int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sValues := make([]string, len(*values))
		for i, s := range *values {
			sValues[i] = strconv.Itoa(s)
		}
		fmt.Fprintln(w, fmt.Sprintf(htmlTemplate, strings.Join(sValues, ",")))
	})
	http.ListenAndServe(":8080", nil)
}
