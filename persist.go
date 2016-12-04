package main

import (
	"fmt"
)

/**
  usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    fmt.Println( usr.HomeDir )
    **/

func indicesPersist(indices chan int) {

	//find slice size ...
	//file to save and so on
	for i := range indices {
		fmt.Printf("index %d\n", i)
	}
}
