package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	
	"ProyectDos/models"
	
)

type Controller struct {
	Automata *models.Automata
	
}

func NewController(automata *models.Automata) *Controller {
	return &Controller{automata}
}

func (c *Controller) LoadAutomataFromFile(filepath string) error {
	// Intenta cargar el archivo especificado.
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	// Decodifica el archivo JSON en un autómata.
	var automata models.Automata
	err = json.Unmarshal(data, &automata)
	if err != nil {
		return err
	}

	// Asigna el autómata cargado al controlador.
	c.Automata = &automata

	return nil
}

func (c *Controller) LoadAutomataFromString(data string) error {
	// Decodifica la cadena de entrada JSON en un autómata.
	var automata models.Automata
	err := json.Unmarshal([]byte(data), &automata)
	if err != nil {
		return err
	}

	// Asigna el autómata cargado al controlador.
	c.Automata = &automata

	return nil
}

func (c *Controller) SaveAutomataToFile(filepath string) error {
	// Codifica el autómata en formato JSON.
	data, err := json.Marshal(c.Automata)
	if err != nil {
		return err
	}

	// Escribe el archivo.
	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) CheckString(str string) (bool, error) {
	if c.Automata == nil {
		return false, errors.New("Automata no definido")
	}

	// Comprueba que el autómata sea completo y lo completa si es necesario.
	err := c.Automata.CheckCompleteness()
	if err != nil {
		return false, err
	}

	// Comprueba si la cadena es aceptada por el autómata.
	return c.Automata.Accept(str), nil
}

