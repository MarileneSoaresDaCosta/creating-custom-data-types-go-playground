package main

import (
	"fmt"
	"playground/organization"
)

func main() {
	p := organization.NewPerson("Jon", "Snow", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))
	err := p.SetTwitterHandler("@king_in_the_north")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Println("An error occurred setting the twitter handler: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.Country())
	println(p.FullName())

}
