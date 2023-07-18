// SOLO HICE LA SUMA
package main

import (
	"fmt"
	"strconv"
)

var enter string
var num1, num2 int

func main() {

	var accion string
	for {

		fmt.Println("Calculadora")
		fmt.Println("Escribe que quieres hacer: sumar, restar, dividir, multiplicar o exit")
		_, err := fmt.Scanln(&accion)
		if err != nil {
			fmt.Println("el error es:\n", err)
			break
		}
		if accion == "exit" {
			fmt.Println("\nPROGRAMA FINALIZADO")
			break
		}

		switch {
		case accion == "sumar":
			suma()
			break
		case accion == "restar":
			resta()
			break
		case accion == "dividir":
			division()
			break
		case accion == "multiplicar":
			multi()
			break
		default:
			fmt.Println("opcion inexistente")
			break
		}

	}
}

func suma() {
	for {
		fmt.Println("Escribe que sumar o exit:")

		//esto es para que pueda escribir y evalua si no es null.
		_, err := fmt.Scanln(&enter)
		if err != nil {
			fmt.Println(err)
		}
		//esto para ver si escribio "exit"
		if enter == "exit" {
			fmt.Println("Saliendo...")
			break
		}

		// esto convierte el string a int y evalua si escribio un numero y no cualquier cosa
		num1, err := strconv.Atoi(enter)
		if err != nil {
			fmt.Println("opcion incorrecta, try again")
			continue
		}

		fmt.Println("escribe el num2 or exit: ")
		_, err = fmt.Scanln(&enter)
		if err != nil {
			fmt.Println(err)
		}
		if enter == "exit" {
			fmt.Println("saliendo..")
			break
		}

		// esto convierte el string a int y evalua si escribio un numero y no cualquier cosa
		num2, err := strconv.Atoi(enter)
		if err != nil {
			fmt.Println("opcion incorrecta , try again")
			continue
		}

		fmt.Println("el resultado es: ", num1+num2)

	}
}
func resta() {
	for {
		fmt.Println("Escribe que sumar o exit:")
	}
}

func division() {
	for {
		fmt.Println("Escribe que sumar o exit:")
	}
}

func multi() {
	for {
		fmt.Println("Escribe que sumar o exit:")
	}
}
