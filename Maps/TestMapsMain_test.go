package main

import "testing"

func TestProcesamiento(t *testing.T) {
	if Procesamiento() != nil {
		t.Errorf("Error en la funcion Procesamiento")
	}
}
