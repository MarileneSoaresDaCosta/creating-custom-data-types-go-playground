package main

import (
	"fmt"
	"playground/organization"
)

func main() {
	p := organization.NewPerson("Jon", "Snow")
	err := p.SetTwitterHandler("@king_in_the_north")
	if err != nil {
		fmt.Println("An error occurred setting the twitter handler: %s\n", err.Error())
	}
	println(p.TwitterHandler())

	println(p.ID())
	println(p.FullName())

}
