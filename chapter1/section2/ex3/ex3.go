//ex3: Experiment to measure the difference in running time between our
// potentially inefficient versions and the one that uses strings.Join.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	echo1()
	fmt.Println(time.Since(start1).Seconds())

	start2 := time.Now()
	echo2()
	fmt.Println(time.Since(start2).Seconds())

	start3 := time.Now()
	echo3()
	fmt.Println(time.Since(start3).Seconds())

	start4 := time.Now()
	echo4()
	fmt.Println(time.Since(start4).Seconds())

}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo4() {
	fmt.Println(os.Args[1:])
}
