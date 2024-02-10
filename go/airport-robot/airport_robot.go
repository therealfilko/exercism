package airportrobot

type Greeter interface {
    LanguageName() string
    Greet(name string) string 
}

type Italian struct {} 

type Portuguese struct {} 

func (n Italian) LanguageName() string {
    return "Italian"
}

func (n Italian) Greet(name string) string {
    return "Ciao " + name
}

func (n Portuguese) LanguageName() string {
    return "Portuguese"
}

func (n Portuguese) Greet(name string) string {
    return "Olá " + name
}

func SayHello(visitorName string, greeter Greeter) string {
    return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(visitorName) + "!"
}
