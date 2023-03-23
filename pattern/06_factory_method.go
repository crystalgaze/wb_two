package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

type Comp interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Comp {
	switch typeName {
	default:
		fmt.Printf("%s несуществующий тип объекта\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	}
}

type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() Comp {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d]\n", s.Type, s.Core, s.Memory)
}

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPersonalComputer() Comp {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (p PersonalComputer) GetType() string {
	return p.Type
}

func (p PersonalComputer) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d] Monitor:[%t]\n", p.Type, p.Core, p.Memory, p.Monitor)
}

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNotebook() Comp {
	return Notebook{
		Type:    NotebookType,
		Core:    4,
		Memory:  8,
		Monitor: true,
	}
}

func (n Notebook) GetType() string {
	return n.Type
}

func (n Notebook) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d] Monitor:[%t]\n", n.Type, n.Core, n.Memory, n.Monitor)
}

var types = []string{PersonalComputerType, NotebookType, ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computer := New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}
