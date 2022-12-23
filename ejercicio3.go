package main

import "fmt"

func main() {
	var minutos int
	var categoria string
	fmt.Println("Ingrese los minutos trabajados porfavor")
	fmt.Scanln(&minutos)
	fmt.Println("Ingrese su categoria porfavor")
	fmt.Scanln(&categoria)
	salario := CalcularSalario(minutos, categoria)

	fmt.Printf("Su salario es: %.2f\n", salario)
}

func CalcularSalario(minutos int, categoria string) float64 {
	var salario float64
	switch categoria {
	case "A":
		salario = float64(minutos/60) * 3000
		salario = salario + salario*0.5
	case "B":
		salario = float64(minutos/60) * 1500
		salario = salario + salario*0.2
	case "C":
		salario = float64(minutos/60) * 1000
	}
	return salario
}
