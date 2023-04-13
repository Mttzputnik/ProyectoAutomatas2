package automataControlador

import (
    "encoding/json"
    "io/ioutil"
    "os"
	"./source/modelo/automataModelo"

)

automataModelo:= automata.NewAutomata()

// Función para cargar el autómata desde un archivo JSON

func LoadAutomataFromJSON(filename string) (*Automata, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var automata Automata
    err = json.Unmarshal(bytes, &automata)
    if err != nil {
        return nil, err
    }

    return &automata, nil
}

