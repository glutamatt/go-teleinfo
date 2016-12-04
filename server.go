package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var htmlTemplate = "<!DOCTYPE html><html><head><link rel=\"stylesheet\"" +
	"href=\"//cdn.jsdelivr.net/chartist.js/latest/chartist.min.css\">" +
	"<script src=\"//cdn.jsdelivr.net/chartist.js/latest/chartist.min.js\"></script>" +
	"</head><body><div class=\"ct-chart ct-minor-sixth\"></div><script>setTimeout(function(){location = ''},5000);" +
	"var data = {series: [[%s]]}; new Chartist.Line('.ct-chart', data, {low: 0,showArea: true, showLine: false, showPoint: false});" +
	"</script></body></html>"

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
