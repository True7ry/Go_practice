package main // Esto es para nombre de la carpeta donde esta guardado.

import "fmt"

/*
Como ejecutar un archivo?: go run
Ejecutarlo para produccion(crea un archivo): go build src/main.go
*/

func escribirEnConsola(mensaje string) {
	fmt.Println(mensaje)

}
func main() {
	//Imprimir en terminal.
	fmt.Println("Terminal")

	//Declaracion de variables    (las variables se declaran especificando su valor
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("el valor de PI es: ", pi)
	fmt.Println("Variable sin especificar = ", pi2)

	//Declaracion de variable enteras

	base := 12 //Poner : antes de declarar una variable, permite declararla sin especificar si es int, bool, string, etc.
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Operadores aritméticos
	sumar := 2 + 2
	restar := 2 - 2
	dividir := 4 / 2
	multiplicar := 2 * 2
	fmt.Println(sumar, restar, dividir, multiplicar)

	//Funciones en Go
	escribirEnConsola("PORQUE SE PONEN ASI")

	// Ciclos(solo existe for)

	condicion := 10

	for i := 0; i >= 10; i++ {
		fmt.Print(i)
	}
	fmt.Printf("\n")

	for condicion > 14 {
		fmt.Print("hola")
		condicion++
	}

	fmt.Printf("\n")
	//	for {//Esto es un for forever
	//		fmt.Print("hola")
	//	}
	fmt.Printf("\n")

	//CONDICIONALES IF
	constans := "vva"

	if constans == "viva" {
		fmt.Println("it's alive")
	} else {
		fmt.Println("no esta viva")
	}

	//Condiciones con Swith(else if no existe)
	switch constans {
	case "viva":
		fmt.Println("it's alive")
	case "Muerta":
		fmt.Println("esta remuerta")
	default:
		fmt.Println("ni idea")
	}

	//Switch sin condicion(como un else if)

	var valor int = 60
	switch {
	case valor == 100:
		fmt.Println("es 100")
	case valor <= 0:
		fmt.Println("es menor que 0")
	case valor > 50:
		fmt.Println("es mayor que 50")
	default:
		fmt.Println("esta entre 1 y 50")
	}

	// Defer, continue y break
	/*
		Defer hace que algo (casi siempre es una funcion) se ejecuta al final del codigo.

		continue y break se utilizan en ciclos "for". Continue sirve para que la instruccion no muere al entrar en una condicion
		y break para cerrar al entrar en una condicion.
	*/
	for i := 0; i <= 10; i++ {

		fmt.Println(i)
		//continue
		if i == 8 {
			fmt.Println("es 8")
			continue // se pasa a la siguiente iteración sin ejecutar el código restante.
		}
		//break
		if i == 8 {
			fmt.Println("close program")
			break
		}
	}

	//Arrays(son inmutables, no puedes cambiar su tamaño)
	var listaa [2]string
	listaa[0] = "miguel"
	listaa[1] = "death"
	fmt.Println(listaa)

	arraysInmutable := [5]string{"pedro", "Analysis", "Missing", "kidnapped", "1983"}

	fmt.Println(arraysInmutable)
	//Slice (son listas comunes que puedes cambiar a gusto, agregando o eliminando)

	var slice []int

	slice = append(slice, 8)

	fmt.Println(slice)

	// "Range" en adicion con "for"

	for i, value := range listaa { // 'i' te retorna el indice y value te retorna los valores en el array
		if i == 0 {
			fmt.Println("\n Range en adiccion con for \n ")
		}
		fmt.Println(i, value)
	}

	for _, value := range listaa { // '_' es para que no te retorne el indice(puedes hacer lo mismo con 'value').
		fmt.Println(value)
	}

	for i := range listaa { // solo tener 'i' hace que te retornen los indices.
		fmt.Println(i)
	}

	// Maps(diccionarios): clave, valor.

	fmt.Println("/nMaps(diccionarios) clave, valor.\n ")

	// Crear un map de tipo string
	//  var mapa map[string]string

	// Inicializar el map
	//mapa = make(map[string]string)

	//Agregar y crear un map

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	// puede hacerlo todo con :=
	nameAge := make(map[string]int)

	nameAge["James"] = 12
	nameAge["Ricardo"] = 32
	fmt.Println(nameAge, temperature)

	//encontrar valor

	real, okk := nameAge["James"] //aqui si existe la llave, pero y si no?
	value, ok := nameAge["www"]   //aqui no existe.

	fmt.Println("\n", real, okk)
	fmt.Println("\n", value, ok) // con 'ok' te suelta un booleano, confirmando si existe o no.

	//CLASES(esta en una funcion aparte)
	fmt.Println("\nCLASES en Go:")
	Clases()

	// Modificadores de acceso en funciones y struct"clases" son: (privado y publica)
	/*
		Al declarar las fuciones o clases en mayusculas, puedes exportarlas y utilizarlas en otros archivos, pero si es en minusculas no puedes.

		func Publica(){
		  fmt.Println("ol")
		}

		func privada(){
		  fmt.Println("funcion privada")

		type ClasePublica struct{
		  name string
		}

		type clasePrivada struct {
		  age int
		}
	*/

	//Structs y punteros

	//Stringers: personalizar el output de Structs
	papelSize := 50
	baconSize := &a

	fmt.Println(papelSize, baconSize)
	//Interfaces y listas de interfaces

	// Que es la concurrencia?
	//Primer contacto con las Goroutines(concurrencias con go)

	fmt.Println("holanda")
	go fmt.Println("Pedregulho") // usualmente no se utiliza, ya que son dificiles de mantener al largo plazo.
	//en su lugar se utiliza los Channels

	//CHANNELS.

	Shannels := make(chan string, 1) // el numero significa cuantos datos maneja

	fmt.Println("World")
	go say("Bye", Shannels)

	fmt.Println(<-Shannels)
}

func say(text string, Shannels chan<- string) { // "<-" antes de "chan" sirve para indicar que recibe informacion, si es al reves(<- tiene que estar al otro lado de "chan") da informacion y no recibe.
	Shannels <- text
	// text <- Shannels(esto significa que da info y no recibe)
}

type car struct {
	// "brand" == "marca".
	brand string
	year  int
}

func Clases() {

	//CLASES EN Go: Structs
	MyNewCar := car{brand: "Ford", year: 2020}
	fmt.Println(MyNewCar)

	//otra manera
	var OtherC car
	OtherC.brand = "Tesla"
	OtherC.year = 2012
	fmt.Println(OtherC)

}

func fmtMasQueUnPrint() {

	const pi float64 = 3.14
	const pi2 = 3.14
	//Paquete "fmt": algo mas que imprimir en consola
	//Printf
	Decimales := "%f"
	Palabras := "%s"
	numEntero := "%d"
	Booleanos := "%t" // si no sabes el tipo que quieres poner, puedes poner "%v"
	fmt.Println(Decimales, Palabras, numEntero, Booleanos)
	fmt.Printf("pedro mide: %f \n", pi)

	// Saber el tipo de dato de una variable SE UTILIZA: %T
	fmt.Printf("El tipo de dato es: %T \n", pi)
}
