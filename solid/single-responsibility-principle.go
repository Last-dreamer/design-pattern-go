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
	// 0644 permission for file
	_ = os.WriteFile(fileName, []byte(strings.Join(j.Entries, p.LineSeparator)), 0644)
}
