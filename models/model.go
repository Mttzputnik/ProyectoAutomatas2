package models

type Automata struct {
    States       []string            `json:"states"`
    Alphabet     []string            `json:"alphabet"`
    Transitions  map[string]map[string]string `json:"transitions"`
    InitialState string              `json:"initial_state"`
    FinalStates  []string            `json:"final_states"`
}
