package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	conuts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr,"dup3: %v\n" , err)
			continue
		}
		for _,line := range strings.Split(string(data),"\n"){
			conuts[line]++
		}
	}
	for line ,n := range conuts{
		if n > 1 {
			fmt.Printf("%d\t %s\n",n,line)
		}
	}
}