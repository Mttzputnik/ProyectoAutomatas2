package controllers

import (
	"encoding/json"
	"errors"
	"os"
	"proyect/models"
)

// función para validar si una cadena de entrada pertenece al lenguaje definido por el autómata:

func (a *Automata) ValidateString(input string) bool {
    currentState := a.startState
    for _, symbol := range input {
        nextState, exists := a.transitions[currentState][string(symbol)]
        if !exists {
            return false
        }
        currentState = nextState
    }
    for _, state := range a.finalStates {
        if currentState == state {
            return true
        }
    }
    return false
}

// función para cargar el autómata desde un archivo de texto o JSON:

func LoadAutomataFromFile(path string) (*Automata, error) {
    // Leer archivo
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Decodificar archivo JSON
    var automata Automata
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&automata)
    if err != nil {
        return nil, err
    }

    // Validar autómata
    if !automata.IsValid() {
        return nil, errors.New("el autómata es inválido")
    }

    return &automata, nil
}

//función para completar un autómata incompleto agregando un estado sumidero:

func (a *Automata) Complete() {
    // Agregar estado sumidero
    a.states = append(a.states, "sumidero")
    for _, symbol := range a.alphabet {
        a.transitions["sumidero"] = map[string]string{symbol: "sumidero"}
    }
    for _, state := range a.states {
        for _, symbol := range a.alphabet {
            _, exists := a.transitions[state][symbol]
            if !exists {
                a.transitions[state][symbol] = "sumidero"
            }
        }
    }
}

