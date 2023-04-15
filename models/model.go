package models

import (
	
	"errors"
	
)

// Automata es la estructura que representa un autómata finito determinístico.
type Automata struct {
	States       []string                     `json:"states"`
	Alphabet     []string                     `json:"alphabet"`
	Transitions  map[string]map[string]string `json:"transitions"`
	InitialState string                       `json:"initial_state"`
	FinalStates  []string                     `json:"final_states"`
	isComplete   bool
}

// NewAutomata crea una nueva instancia del autómata.
func NewAutomata() *Automata {
	return &Automata{
		Transitions: make(map[string]map[string]string),
	}
}

func (a *Automata) Complete() {
	// Agregar estado sumidero
	a.States = append(a.States, "sumidero")
	for _, symbol := range a.Alphabet {
		a.Transitions["sumidero"] = map[string]string{symbol: "sumidero"}
	}
	for _, state := range a.States {
		for _, symbol := range a.Alphabet {
			_, exists := a.Transitions[state][symbol]
			if !exists {
				a.Transitions[state][symbol] = "sumidero"
			}
		}
	}
	a.isComplete = true
}

func (a *Automata) CheckCompleteness() error {
	if !a.isComplete {
		return errors.New("Automata is incomplete")
	}
	for _, state := range a.States {
		if _, ok := a.Transitions[state][""]; !ok {
			return errors.New("Automata is incomplete")
		}
		for _, sym := range a.Alphabet {
			if _, ok := a.Transitions[state][sym]; !ok {
				return errors.New("Automata is incomplete")
			}
		}
	}
	return nil
}

// Accept verifica si la cadena de entrada es aceptada por el autómata.
func (a *Automata) Accept(cadena string) bool {
	estadoActual := a.InitialState
	for _, simbolo := range cadena {
		if _, ok := a.Transitions[estadoActual][string(simbolo)]; !ok {
			return false
		}
		estadoActual = a.Transitions[estadoActual][string(simbolo)]
	}
	for _, estadoFinal := range a.FinalStates {
		if estadoActual == estadoFinal {
			return true
		}
	}
	return false
}
