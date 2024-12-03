package designpatterns

type Person struct {
	// address
	StreetAddress, PostCode, City string

	//job informations
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(StreetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = StreetAddress
	return it
}

func (it *PersonAddressBuilder) In(City string) *PersonAddressBuilder {
	it.person.City = City
	return it
}

func (it *PersonAddressBuilder) WithPostCode(postCode string) *PersonAddressBuilder {
	it.person.PostCode = postCode
	return it
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(CompanyName string) *PersonJobBuilder {
	pjb.person.CompanyName = CompanyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(Position string) *PersonJobBuilder {
	pjb.person.Position = Position
	return pjb
}
