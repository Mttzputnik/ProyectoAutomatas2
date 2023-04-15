package views

import (
	"fmt"
	"os"

	"ProyectDos/controllers"
)

type View struct {
	Controller *controllers.Controller
}

func NewView(controller *controllers.Controller) *View {
	return &View{controller}
}

func (v *View) GetInput() {
	var input string
	fmt.Println("Ingrese una cadena para verificar:")
	fmt.Scanln(&input)

	isAccepted, err := v.Controller.CheckString(input)
	if err != nil {
		fmt.Println("Error al verificar la cadena:", err)
	} else if isAccepted {
		fmt.Println("La cadena", input, "es aceptada por el autómata")
	} else {
		fmt.Println("La cadena", input, "no es aceptada por el autómata")
	}

	
}

func (v *View) Run() {
	for {
		var input string
		fmt.Println("¿Qué desea hacer?")
		fmt.Println("1. Cargar autómata desde archivo")
		fmt.Println("2. Cargar autómata desde entrada manual")
		fmt.Println("3. Guardar autómata en archivo")
		fmt.Println("4. Verificar cadena")
		fmt.Println("5. Salir")
		fmt.Scanln(&input)

		switch input {
		case "1":
			fmt.Println("Ingrese la ruta del archivo:")
			fmt.Scanln(&input)
			err := v.Controller.LoadAutomataFromFile(input)
			if err != nil {
				fmt.Println("Error al cargar el archivo:", err)
			} else {
				fmt.Println("Autómata cargado correctamente")
			}
		case "2":
			fmt.Println("Ingrese el autómata en formato JSON:")
			fmt.Fscanln(os.Stdin, &input)
			err := v.Controller.LoadAutomataFromString(input)
			if err != nil {
				fmt.Println("Error al cargar el autómata:", err)
			} else {
				fmt.Println("Autómata cargado correctamente")
			}
		case "3":
			fmt.Println("Ingrese la ruta del archivo:")
			fmt.Scanln(&input)
			err := v.Controller.SaveAutomataToFile(input)
			if err != nil {
				fmt.Println("Error al guardar el archivo:", err)
			} else {
				fmt.Println("Autómata guardado correctamente")
			}
		case "4":
			v.GetInput()
		case "5":
			fmt.Println("Saliendo...")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida")
		}
	}
}
