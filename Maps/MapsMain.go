package main

import (
	"errors"
	"fmt"
)

var ErrInvalidVehicle = errors.New("invalid vehicle")


func main() {
	Procesamiento()
}

func Procesamiento() error {
	m := map[string]int{
		"hola": 12124,
	}
	fmt.Println(len(m))
	valor := m["hola"]
	fmt.Println(valor)

	valor2, error := m["hola"]

	fmt.Println(valor2,error)

	if error == false {
		panic("Error por false!")
		return errors.New("Error interno")
	}

	for clave,valor := range m {
		fmt.Println("clave:",clave,"- valor:",valor)
	}
	v:="hola"
	b:=2
	fmt.Println(fmt.Sprintf("%s-%d", v, b))

	if err := funcionTest(); err != nil {
		//panic("Error en funcion:",err)
		fmt.Println(err)
		//return err
	}

	err := funcionTest2()
	fmt.Println(err)
	//if err := funcionTest2(); err != nil {
	//	fmt.Println(err)
	//}
	return nil
}

func funcionTest() error{
	return ErrInvalidVehicle
}

func funcionTest2() InvalidStatusError{
	return NewInvalidStatusError("Error inyectado")
}