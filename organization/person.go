package organization

import (
	"errors"
	"fmt"
	"strings"
)

// this creates an alias; underneath, it is still a string
//type TwitterHandler = string

type Handler struct {
	handle string
	name   string
}

// alternatively, we can simply us a type declaration
type TwitterHandler string

// this new method - RedirectUrl -  would not be possible to define if we were using an alias
func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	// define here the functions I want
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}
type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States of America"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

type Name struct {
	firstName string
	lastName  string
}
type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
}

// go has no constructor, so to keep control of new obj, we use the following idiom
func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			firstName: firstName,
			lastName:  lastName,
		},
		Citizen: citizen,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

// this is no longer necessary when we embedded the interface Identifiable directly on Person
//func (p *Person) ID() string {
//	return "12345"
//}

func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
}

// when editing state with a custom type, we must use a pointer
// we need to use a pointer-base receiver
// because a function makes a copy of each argument value, if a func needs to
// update a var, we must pass the address of the variable using a pointer
// this is also important when large data is being passed
// *** by convention, then we make all methods use the pointer receiver, even the read-only ones

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 { // we want the user to be able to create an empty string handler
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
