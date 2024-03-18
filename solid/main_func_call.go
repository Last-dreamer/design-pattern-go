package solid

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
	// apple := Product{
	// 	Name:  "Apple",
	// 	Color: green,
	// 	Size:  small,
	// }
	// house := Product{
	// 	Name:  "House",
	// 	Color: blue,
	// 	Size:  large,
	// }
	// tree := Product{
	// 	Name:  "Tree",
	// 	Color: green,
	// 	Size:  large,
	// }

	// products := []Product{apple, house, tree}
	// fmt.Printf("Green Old Products :\n")

	// f := Filter{}
	// for _, v := range f.FilterByColor(products, green) {
	// 	fmt.Printf(" - %s is green\n", v.Name)
	// }

	// fmt.Printf("Green Product (New): \n")
	// greenSpec := ColorSpicification{green}
	// bf := BetterFilter{}
	// for _, v := range bf.Filter(products, &greenSpec) {
	// 	fmt.Printf(" - %s is green\n", v.Name)
	// }

	// fmt.Printf("Large Green Product: \n")
	// largeSpec := SizeSpicification{large}
	// lgSpec := AndSpecification{&largeSpec, &greenSpec}
	// for _, v := range bf.Filter(products, &lgSpec) {
	// 	fmt.Printf(" - %s is green\n", v.Name)
	// }

	//! liskov substitution principle

	// rc := &Rectangle{2, 3}
	// UseIt(rc) // 20, 20

	// breaking the liskov substitution principle
	// sq := NewSquare(5)
	// UseIt(sq) // 20, 50

	// // frist solution
	// sq2 := &Square2{5}
	// UseIt(sq2.Rectangle()) // 50, 50

	//! interface segregation principle
	// ofp := OldFunctionPrinter{}
	// ofp.Scan()

}
