package solid

import (
	"log"
)

type Documents struct{}

type Machine interface {
	Print(d Documents)
	Fax(d Documents)
	Scan(d Documents)
}

type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Documents) {
	panic("not implemented") // TODO: Implement
}

func (m MultiFunctionPrinter) Fax(d Documents) {
	panic("not implemented") // TODO: Implement
}

func (m MultiFunctionPrinter) Scan(d Documents) {
	log.Println("scaning.......")
	// panic("not implemented") // TODO: Implement
}

type OldFunctionPrinter struct{}

func (o OldFunctionPrinter) Print(d Documents) {
	panic("not implemented") // TODO: Implement
}

// Defricated: ...
func (o OldFunctionPrinter) Fax(d Documents) {
	panic("operation no supported yet")
}

// Defrecated: ...
func (o OldFunctionPrinter) Scan(d Documents) {
	panic("operation no supported yet")
}

// ISP :=  try to break up the interface into seperate part that the people would difinitly need
// better approach: split into several interfaces

type Printer interface {
	Print(d Documents)
}
type Scanner interface {
	Scan(d Documents)
}

// eg
type MyPrinter struct{}

func (m MyPrinter) Print(d Documents) {}

// another example
type PhotoCopier struct{}

func (p PhotoCopier) Print(d Documents) {}
func (p PhotoCopier) Scan(d Documents) {
	log.Println("testing scanning interface ..")
}

//  to compose more interfaces into one

type MulitFunctionDevice interface {
	Printer
	Scanner
	// if have another functionality
	// Fax
}

// decorator pattern
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
	// may be other interfaces also
}

func (m MultiFunctionMachine) Print(d Documents) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Documents) {
	// log.Println("tesitng if working or not ")
	m.scanner.Scan(d)
	// fmt.Printf("m.Scanner: just scanning .. %v\n", m.Scanner)

}
