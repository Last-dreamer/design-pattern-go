package solid

type Documents struct{}

type Machine interface {
	Print(d Documents)
	Fax(d Documents)
	Scan(d Documents)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Documents) {
	panic("not implemented") // TODO: Implement
}

func (m MultiFunctionPrinter) Fax(d Documents) {
	panic("not implemented") // TODO: Implement
}

func (m MultiFunctionPrinter) Scan(d Documents) {
	panic("not implemented") // TODO: Implement
}

type OldFunctionPrinter struct{}

func (o OldFunctionPrinter) Print(d Documents) {
	panic("not implemented") // TODO: Implement
}

// Defrecated: ...
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
func (p PhotoCopier) Scan(d Documents)  {}

//  to compose more interfaces into one

type MulitFunctionDevice interface {
	Printer
	Scanner
	// if have another functionality
	// Fax
}

// decorator pattern
type MultiFunctionMachine struct {
	Printer Printer
	Scanner Scanner
}

func (m MultiFunctionMachine) Print(d Documents) {
	m.Printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Documents) {
	m.Scanner.Scan(d)
}
