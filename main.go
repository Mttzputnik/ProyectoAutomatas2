package main

import (
	

	"ProyectDos/views"
	"ProyectDos/controllers"
	"ProyectDos/models"
)
func main() {
    // Inicializar el modelo
    automata := &models.Automata{
        States:  []string{"q0", "q1", "q2"},
        Alphabet: []string{"a", "b"},
        Transitions: map[string]map[string]string{
            "q0": {
                "a": "q1",
                "b": "q2",
            },
            "q1": {
                "a": "q0",
                "b": "q2",
            },
            "q2": {
                "a": "q2",
                "b": "q1",
            },
        },
    }
	// Inicializar el controlador
controller := controllers.NewController(automata)

// Inicializar la vista
view := views.NewView(controller)

// Ejecutar la vista
view.Run()
}