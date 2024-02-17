package solid

// open for extension but closed for modification
// enterprise pattern := Specification
type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	Name  string
	Color Color
	Size  Size
}

type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {

	result := make([]*Product, 0)
	for i, v := range products {
		if v.Color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// creating an interface for specifiction
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpicification struct {
	Color Color
}

func (c *ColorSpicification) IsSatisfied(p *Product) bool {
	return c.Color == p.Color
}

type SizeSpicification struct {
	Size Size
}

func (s *SizeSpicification) IsSatisfied(p *Product) bool {
	return s.Size == p.Size
}

type BetterFilter struct{}

func (b *BetterFilter) Filter(p []Product, spec Specification) []*Product {

	result := make([]*Product, 0)
	for i, v := range p {
		if spec.IsSatisfied(&v) {
			result = append(result, &p[i])
		}
	}
	return result
}

// composite pattern
// --  partitioning design pattern
// --  it's a combinator which combine two diff specification

type AndSpecification struct {
	first, second Specification
}

func (a *AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}
