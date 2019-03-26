package Index

import (
	"errors"
	"fmt"
)

type Punto struct {
	valor1 string
	valor2 string
}

func (p *Punto) funcion1() int {
	fmt.Println("Entró en la función del objeto")
	return 1
}

type PuntoExtra struct {
	Punto
	valor3 string
}

type Union interface {
	funcion1() int
}

func unaFuncion(p Union) int{
	fmt.Println("valor dentro de la funcion")
	return p.funcion1()
}

func main() {
	p1:=Punto{"",""}
	p1.funcion1()
	unaFuncion(&p1)

	p2:=PuntoExtra{Punto{"",""},""}
	p2.funcion1()
	unaFuncion(&p2)

	fmt.Println("Hola!!")
	if error := testeo(); error != nil {
		fmt.Println("Error:", error)
		panic("Panic!!")
	}
	fmt.Println("Adios!!")
}

func testeo() error {
	if b:=true; b==true {
		fmt.Println("retorno 1")
		return errors.New("Error retornado por true")
	}
	fmt.Println("retorno 2")
	return nil
}