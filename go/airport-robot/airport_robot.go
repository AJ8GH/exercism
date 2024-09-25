package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return g.Greet(name)
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "I can speak Italian: "
}

func (i Italian) Greet(name string) string {
	return i.LanguageName() + fmt.Sprintf("Ciao %v!", name)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "I can speak Portuguese: "
}

func (p Portuguese) Greet(name string) string {
	return p.LanguageName() + fmt.Sprintf("Olá %v!", name)
}
