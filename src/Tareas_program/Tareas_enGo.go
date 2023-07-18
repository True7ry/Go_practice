package main

/*
Al poner espacios entre los strings, se rompe. (ni idea de como arreglarlo).
*/
import (
	"fmt"
)

func main() {
	var gbv string
	var input string = "hola n da"
	var slice []string
	var find bool = false
	for {

		fmt.Println("Tasks Application in Go")
		fmt.Println("Que quieres hacer? Guardar, Borrar o Ver? ")
		fmt.Println("'exit' para salir.")
		fmt.Scan(&gbv)

		if gbv == "exit" || gbv == "Exit" {
			fmt.Println("Programa finalizado")
			break
		}

		switch {
		case gbv == "Guardar" || gbv == "guardar":
			for {
				fmt.Println("Escribe la tarea(o exit):")
				_, err := fmt.Scanln(&input)

				if err != nil {
					fmt.Println("Error al leer la entrada:", err)
					return
				}

				if input == "exit" {
					fmt.Println("saliendo")
					break
				}

				slice = append(slice, input)
			}
		case gbv == "Borrar" || gbv == "borrar":
			for {
				find = false
				fmt.Println("Que tarea quieres borrar?")
				_, err := fmt.Scanln(&input)

				if err != nil {
					fmt.Println("Error: ", err)
					break
				}

				if input == "exit" {
					fmt.Println("saliendo...")
					break
				}
				for _, nombre := range slice {
					if nombre == input {
						find = true
						break
					}
				}
				if find {
					fmt.Println("No se elimino nada, ya que no encontre la funcion para eliminar un elemento de un slice sin crear otro slice")
				} else {
					fmt.Println("elemento no encontrado")
				}
			}
		case gbv == "Ver" || gbv == "ver":
			fmt.Println("Tus Tareas son:\n ")
			for _, elemento := range slice {
				fmt.Println(elemento)
			}
			fmt.Println("esas son todas tus tareas \n ")

		default:
			fmt.Println("opcion incorrecta")
		}
	}
}
