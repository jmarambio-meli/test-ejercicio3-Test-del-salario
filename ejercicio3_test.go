package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularSalario(t *testing.T) {

	minutos1 := 60
	minutos2 := 120
	minutos3 := 180

	categoria1 := "A"
	categoria2 := "B"
	categoria3 := "C"

	var salarioesperado1 float64
	var salarioesperado2 float64
	var salarioesperado3 float64

	salarioesperado1 = float64(minutos1/60) * 3000
	salarioesperado1 = salarioesperado1 + salarioesperado1*0.5

	salarioesperado2 = float64(minutos2/60) * 1500
	salarioesperado2 = salarioesperado2 + salarioesperado2*0.2

	salarioesperado3 = float64(minutos3/60) * 1000

	resultado1 := CalcularSalario(minutos1, categoria1)
	resultado2 := CalcularSalario(minutos2, categoria2)
	resultado3 := CalcularSalario(minutos3, categoria3)

	assert.Equal(t, salarioesperado1, resultado1, "Salario incorrecto para categoria A")
	assert.Equal(t, salarioesperado2, resultado2, "Salario incorrecto para categoria B")
	assert.Equal(t, salarioesperado3, resultado3, "Salario incorrecto para categoria C")
}
