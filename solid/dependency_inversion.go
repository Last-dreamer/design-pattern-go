package solid

import "log"

// Depency Inversion Priciple
// HLM should not be dependent upon LLM
// Both should be dependented upon abstraction

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	Name string
	// etc...
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low level modules like storing these in the db....
type RelationshipBrowser interface {
	findAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) findAllChildrenOf(name string) []*Person {
	relations := make([]*Person, 0)

	for i, rel := range r.relations {
		if rel.relationship == Parent && rel.from.Name == name {
			relations = append(relations, r.relations[i].to)
		}
	}

	return relations
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
}

// high level module
type Research struct {
	// Break DIP
	// relationships Relationships
	// with DIP
	browser RelationshipBrowser
}

func (r *Research) Invistigate() {
	//! breaking DIP
	// relations := r.relationships.relations

	// for _, rel := range relations {
	// 	if rel.from.Name == "Jhon" && rel.relationship == Parent {
	// 		log.Println("Jhon has the child called , ", rel.to.Name)
	// 	}
	// }

	// ! with DIP
	for _, p := range r.browser.findAllChildrenOf("Jhon") {
		log.Println("Jhon has the child called", p.Name)

	}

}
