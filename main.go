package main

import (
	"design-pattern-go/solid"
	"fmt"
)

func main() {

	//single responsibility principle
	j := solid.Journal{}
	j.AddEntry("i wannan to go ")
	j.AddEntry("i wannan to go to the mall")
	fmt.Println(j.String())

	//
	p := solid.Persistance{LineSeparator: "\n"}
	p.SaveToFile(&j, "testing.txt")

}
