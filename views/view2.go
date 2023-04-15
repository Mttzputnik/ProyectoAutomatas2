package views

import (
    "fmt"
    "fyne.io/fyne"
    "fyne.io/fyne/layout"
    "fyne.io/fyne/widget"
    
    "ProyectDos/controllers"
)

type View struct {
    Controller *controllers.Controller
}

func NewView(controller *controllers.Controller) *View {
    return &View{controller}
}

func (v *View) buildMainUI() fyne.CanvasObject {
    // Label to show the loaded automaton
    stateLabel := widget.NewLabel("Autómata no cargado")

    // Textfield for input string
    inputEntry := widget.NewEntry()

    // Button to perform verification
    verifyButton := widget.NewButton("Verificar", func() {
        isAccepted, err := v.Controller.CheckString(inputEntry.Text)
        if err != nil {
            stateLabel.SetText("Error al verificar la cadena: " + err.Error())
        } else if isAccepted {
            stateLabel.SetText("La cadena " + inputEntry.Text + " es aceptada por el autómata")
        } else {
            stateLabel.SetText("La cadena " + inputEntry.Text + " no es aceptada por el autómata")
        }
    })

    // Create a container with all the UI elements
    container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
        stateLabel,
        widget.NewLabel("Ingrese una cadena para verificar:"),
        inputEntry,
        verifyButton,
    )

    return container
}

func (v *View) Run() {
    // Create a new Fyne application
    app := fyne.NewApp()

    // Create a new window
    window := app.NewWindow("Verificador de autómatas")

    // Add the UI elements to the main content of the window
    window.SetContent(v.buildMainUI())

    // Show the window
    window.ShowAndRun()
}
