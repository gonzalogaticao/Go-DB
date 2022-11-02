package main

import (
	"fmt"
)

func main() {
	//1.-Recibir al participante.
	fmt.Println("Inicio del cuestionario")

	var day int
	fmt.Print("Ingresar dia: ")
	fmt.Scanln(&day)

	var month int
	fmt.Print("Ingresar mes: ")
	fmt.Scanln(&month)

	var year int
	fmt.Print("Ingresar anio: ")
	fmt.Scanln(&year)

	var rut string
	fmt.Print("Ingresar rut: ")
	fmt.Scanln(&rut)

	var name string
	fmt.Print("Ingresar nombre: ")
	fmt.Scanln(&name)

	var lastName string
	fmt.Print("Ingresar apellido: ")
	fmt.Scanln(&lastName)

	var age int
	fmt.Print("Ingresar edad: ")
	fmt.Scanln(&age)

	/*--- Guardar informacion en BD.--- */

	//2.-Imprimir Instrucicciones.
	//3.-Realizar preguntas.
	//4.-Obtener respuestas.
	var answers [10]int
	var i int
	var num int

	for i = 0; i < 10; i++ {
		fmt.Printf("Respuesta de pregunta (%v): ", i+1)
		fmt.Scanln(&num)
		if num >= 0 && num <= 5 {
			answers[i] = num
		} else {
			i--
		}
	}

	//5.-Calcular resultados.
	var result int
	for i = 0; i < 10; i++ {
		result = result + answers[i]
	}
	/*--- Guardar informacion en BD.--- */

	fmt.Printf("El resultado es: %v\n", result)
	if result < 20 {
		fmt.Printf("Participante %v %v se encuentra en buen estado.\n", name, lastName)
	}
	if result >= 20 && result <= 30 {
		fmt.Printf("Participante %v %v se encuentra en condicion estable.\n", name, lastName)
	}
	if result > 30 && result <= 40 {
		fmt.Printf("Participante %v %v requiere atencion.\n", name, lastName)
	}
	if result > 40 {
		fmt.Printf("Participante %v %v se le diagnostica depresion.\n", name, lastName)
	}

	//6.-Cerrar y finalizar.
	fmt.Println("FIN.")
}
