package main

import (
	"flag"
	"fmt"
)

func main() {
	var i int
	var s string
	flag.IntVar(&i, "i", 110, "INT类型")
	flag.StringVar(&s, "s", "", "STRING类型")
	flag.Parse()
	fmt.Printf("i = %v, s = %v", i, s)
}
