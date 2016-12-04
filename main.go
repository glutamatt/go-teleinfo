package main

func main() {
	valuesCount := 100
	values := []int{}

	fileStrings := make(chan string)
	measures := make(chan int)
	indices := make(chan int)
	go readFile("/dev/ttyS0", fileStrings)
	go parseLines(fileStrings, measures, indices)
	go indicesPersist(indices)
	go runHTTPServer(&values)

	for m := range measures {
		values = append(values, m)
		if len(values) > valuesCount {
			values = values[1:valuesCount]
		}
	}
}
