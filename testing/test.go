package testing

type Person struct {
	Age int
}

func (p *Person) isChild() bool {
	return p.Age < 18
}
