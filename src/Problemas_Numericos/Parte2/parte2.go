package main

import "fmt"

func main() {
	/*
		Parte 2:
		Criar um programa na linguagem Go que trabalhe com o operador % e for.
		- Voce e seus colegas querem cirar em formato de código uma antiga brincadeira:
		Ao falar os números sempre que aparecer um múltiplo de 3, o participante deve falar "Pin" e nos múltiplos de 5 o participante deve falar "Pan". Então, seu programa deve imprimir números de 1 a 100 e nos casos citados, não devem aparecer os números, mas sim o que o participante deve falar.
		obs.: quando for múltiplo de 3 e de 5 deve falar "Pin Pan"
	*/
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin Pan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
