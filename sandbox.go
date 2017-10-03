package main

import (
	"regexp"

	"strings"
	"time"
	"fmt"
)



func tryString() {
	input := "frexit, brexit, grexit, nexit"
	strings.Contains(input, "frexit")
	strings.Contains(input, "brexit")
	strings.Contains(input, "grexit")
	strings.Contains(input, "nexit")
}

var pattern = ".*"//"(br|n|gr)exit"
var rgx, _ = regexp.Compile(pattern)

func tryRegex() {

	input := "frexit, brexit, grexit, nexit"

	rgx.FindAllString(input, -1)
}

func track(start time.Time, name string) {
	duration := time.Since(start)
	fmt.Printf("Duration for %s: %s", name, duration)
}

func main() {

	defer track(time.Now(), "Evaluation")

	for i := 0; i <1000000 ; i++ {
		tryRegex()
		//tryString()
	}


}
