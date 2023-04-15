package main

import (
    "log"

    "proyect/controllers"
    "proyect/models"
    "proyect/views"
)

func main() {


    // Creamos el controlador
    automata, err := controllers.LoadAutomataFromFile("automata.json")
    if err != nil {
        log.Fatal(err)
    }

    // Completamos el autómata si es necesario
    CompleteAutomata(automata)

    // Creamos la vista y le pasamos el controlador
    view := views.NewView(controller)

    // Mostramos la vista
    if err := view.Show(); err != nil {
        log.Fatal(err)
    }
}