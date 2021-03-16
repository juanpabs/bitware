package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/bitware/e-base/base"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Println("Buenas Por favor ingrese los datos de la persona a registrar:")
	fmt.Printf("Nombre:")
	nombre, _ := reader.ReadString('\n')
	nombre = nombre[:len(nombre)-1]

	fmt.Printf("Edad (años):")
	edadstr, _ := reader.ReadString('\n')
	edad, _ := strconv.ParseUint(edadstr[:len(edadstr)-1], 10, 64)

	fmt.Printf("Sexo (M/H):")
	sexo, _ := reader.ReadString('\n')
	sexo = sexo[:len(sexo)-1]

	fmt.Printf("Peso (kg):")
	pesostr, _ := reader.ReadString('\n')
	peso, _ := strconv.ParseFloat(pesostr[:len(pesostr)-1], 32)

	fmt.Printf("Altura (m):")
	alturastr, _ := reader.ReadString('\n')
	altura, _ := strconv.ParseFloat(alturastr[:len(alturastr)-1], 32)

	persona := base.NuevaPersona(nombre, uint(edad), sexo, float32(peso), float32(altura))

	fmt.Print("\033[H\033[2J")

	IMC := persona.CalcularIMC()
	if IMC < 0 {
		fmt.Println(persona.GetNombre(), "está por debajo de su peso ideal")
	} else if IMC > 0 {
		fmt.Println(persona.GetNombre(), "está por arriba de su peso ideal")
	} else {
		fmt.Println(persona.GetNombre(), "está dentro de su peso ideal")
	}

	if persona.GetEdad() < 18 {
		fmt.Println(persona.GetNombre(), "es menor de edad")
	} else {
		fmt.Println(persona.GetNombre(), "es mayor de edad")
	}

	fmt.Println(persona.ToString())
}
