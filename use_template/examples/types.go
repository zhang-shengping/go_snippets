package examples

import (
	"fmt"
	"time"
)

type Todo struct {
	Name        string
	Description string
}

type Person struct {
	Name   string
	Emails []string
}

type Product struct {
	Name            string
	Price           float32
	Description     string
	ShippingDate    time.Time
	Sale            bool
	SaleImagePath   []string
	MyFunc          func(string, string) string
	ShippingOptions []string
}

// We can call methods into templates.
// The method should have only one or two return values
// If two values are returned, the second one must be of type error.
func (p Product) Foo() string {
	return "FOO"
}

func (p Product) Bar(test string) string {
	return fmt.Sprintf("Bar: %s", test)
}
