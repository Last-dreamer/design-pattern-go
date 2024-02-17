package solid

import (
	"fmt"
	"os"
	"strings"
)

var entriesCount = 0

type Journal struct {
	Entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.Entries, "\n")
}

// return the position of the entry
func (j *Journal) AddEntry(text string) int {
	entriesCount++
	entry := fmt.Sprintf("%d: %s", entriesCount, text)
	j.Entries = append(j.Entries, entry)
	return entriesCount
}

func (j *Journal) RemoveEntry(position int) {
	// ...
}

// sepation of concerns
var lineSeperator = "\n"

func Save(j *Journal, fileName string) {
	// 0644 permission for file
	_ = os.WriteFile(fileName, []byte(strings.Join(j.Entries, lineSeperator)), 0644)
}

type Persistance struct {
	LineSeparator string
}

func (p *Persistance) SaveToFile(j *Journal, fileName string) {
	_ = os.WriteFile(fileName, []byte(strings.Join(j.Entries, p.LineSeparator)), 0644)
}

func MainFunctionCall() {

	//! single responsibility principle
	// j := Journal{}
	// j.AddEntry("i wannan to go ")
	// j.AddEntry("i wannan to go to the mall")
	// fmt.Println(j.String())
	//
	// p := Persistance{LineSeparator: "\n"}
	// p.SaveToFile(&j, "testing.txt")

	//! open-closed principle
	apple := Product{
		Name:  "Apple",
		Color: green,
		Size:  small,
	}
	house := Product{
		Name:  "House",
		Color: blue,
		Size:  large,
	}
	tree := Product{
		Name:  "Tree",
		Color: green,
		Size:  large,
	}

	products := []Product{apple, house, tree}
	fmt.Printf("Green Old Products :\n")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.Name)
	}

	fmt.Printf("Green Product (New): \n")
	greenSpec := ColorSpicification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, &greenSpec) {
		fmt.Printf(" - %s is green\n", v.Name)
	}

	fmt.Printf("Large Green Product: \n")
	largeSpec := SizeSpicification{large}
	lgSpec := AndSpecification{&largeSpec, &greenSpec}
	for _, v := range bf.Filter(products, &lgSpec) {
		fmt.Printf(" - %s is green\n", v.Name)
	}
}
