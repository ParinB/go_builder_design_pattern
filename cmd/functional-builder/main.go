package main

import "fmt"

type  Person struct {
	name,position string
}
// Takes A Person And Does Something To It
type personMod func(*Person)

type PersonBuilder struct {
	actions  []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}
// in the builder specify that a person is called something
func (b *PersonBuilder) Build() *Person{
	p := Person{}
	for _,a := range b.actions {
		a(&p)
	}
	return  &p
}
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder{
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return  b
}
func main() {
	b := PersonBuilder{}
	p := b.Called("Parin").WorksAsA("Develeper").Build()
	fmt.Println(*p)
}