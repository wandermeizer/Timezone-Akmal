package main

import (
	"fmt"
	"time"
)

// main function
func main() {

	// with the help of time.Now()
	// store the local time
	t := time.Now()

	// print location and local time
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}

	// Asia/Singapore
	fmt.Println("Location : ", location, " Time : ", t.In(location))
}
