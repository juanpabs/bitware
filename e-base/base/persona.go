package base

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	hombre = "H"
	mujer  = "M"

	minPesoIdealHombre = 20
	maxPesoIdealHombre = 25

	minPesoIdealMujer = 19
	maxPesoIdealMujer = 24
)

type Persona struct {
	nombre string
	edad   uint
	NSS    string
	sexo   string
	peso   float32
	altura float32
}

func NuevaPersona(nombre string, edad uint, sexo string, peso float32, altura float32) *Persona {
	return &Persona{
		nombre: nombre,
		edad:   edad,
		NSS:    generaNSS(),
		sexo:   obtenerSexo(sexo),
		peso:   peso,
		altura: altura,
	}
}

func (p *Persona) CalcularIMC() int {
	pesoIdeal := p.peso / float32(math.Pow(float64(p.altura), 2))

	if p.sexo == hombre {
		return hombreIMC(pesoIdeal)
	} else {
		return mujerIMC(pesoIdeal)
	}

}

func hombreIMC(pesoIdeal float32) int {
	if pesoIdeal < minPesoIdealHombre {
		return -1
	} else if pesoIdeal > maxPesoIdealHombre {
		return 1
	} else {
		return 0
	}
}

func mujerIMC(pesoIdeal float32) int {
	if pesoIdeal < minPesoIdealMujer {
		return -1
	} else if pesoIdeal > maxPesoIdealMujer {
		return 1
	} else {
		return 0
	}
}

func (p *Persona) esMayorDeEdad() bool {
	if p.edad >= 18 {
		return true
	} else {
		return false
	}
}

func comprobarSexo(sexo string) bool {
	if sexo != hombre && sexo != mujer {
		return false
	} else {
		return true
	}
}

func (p *Persona) ToString() string {
	return fmt.Sprintf("Nombre: %s\n Edad: %d \n NSS: %s \n Sexo: %s \n Peso: %f \n Altura: %f",
		p.nombre, p.edad, p.NSS, p.sexo, p.peso, p.altura)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generaNSS() string {
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (p *Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) GetSexo() string {
	return p.sexo
}

func (p *Persona) GetNSS() string {
	return p.NSS
}

func (p *Persona) GetEdad() uint {
	return p.edad
}

func (p *Persona) GetPeso() float32 {
	return p.peso
}

func (p *Persona) GetAltura() float32 {
	return p.altura
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona) SetSexo(sexo string) {
	p.sexo = obtenerSexo(sexo)
}

func (p *Persona) SetPeso(peso float32) {
	p.peso = peso
}

func (p *Persona) SetAltura(altura float32) {
	p.altura = altura
}

func (p *Persona) SetEdad(edad uint) {
	p.edad = edad
}

func obtenerSexo(sexo string) string {
	if comprobarSexo(sexo) {
		return sexo
	} else {
		return hombre
	}
}
