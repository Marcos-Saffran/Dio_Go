package main

import "fmt"

func main() {
	var pontoDeEbulicaoDaAguaEmKelvin int = 373
	var pontoDeEbulicaoDaAguaEmGrausCelsius int = pontoDeEbulicaoDaAguaEmKelvin - 273

	fmt.Println("Temperatura de ebulição da água em Kelvin é igual a", pontoDeEbulicaoDaAguaEmKelvin, "K e a temperatura de ebulição da água em Celsius é igual a", pontoDeEbulicaoDaAguaEmGrausCelsius, "°C")

}
