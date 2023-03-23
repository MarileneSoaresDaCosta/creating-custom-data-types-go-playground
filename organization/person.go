package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	// define here the functions I want
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHandler string
}

// go has no constructor, so to keep control of new obj, we use the following idiom
func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) ID() string {
	return "12345"
}

// when editing state with a custom type, we must use a pointer
// we need to use a pointer-base receiver
// because a function makes a copy of each argument value, if a func needs to
// update a var, we must pass the addres of the variable using a pointer
// this is also important when large data is being passed
// *** by convention, then we make all methods use the pointer receiver, even the read only ones

func (p *Person) SetTwitterHandler(handler string) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() string {
	return p.twitterHandler
}
