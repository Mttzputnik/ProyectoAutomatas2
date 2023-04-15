package views


import (
    "fyne.io/fyne/app"
    "fyne.io/fyne/widget"
    "fyne.io/fyne/layout"
    "fyne.io/fyne/dialog"
)

func Diseñar() {
	// Crear ventana
	app := fyne.NewApp()
	window := app.NewWindow("Validador de Cadenas")
	window.Resize(fyne.NewSize(400, 300))

	// Crear widgets
	label := widget.NewLabel("Ingrese una cadena:")
	input := widget.NewEntry()
	button := widget.NewButton("Validar", func() {
		if automata.ValidateString(input.Text) {
			dialog.ShowInformation("Cadena válida", "La cadena ingresada pertenece al lenguaje.")
		} else {
			dialog.ShowInformation("Cadena inválida", "La cadena ingresada no pertenece al lenguaje.")
		}
	})

	// Crear layout
	content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		label,
		input,
		button,
	)

	// Setear contenido
	window.setContent(content)

	// Show window and run app
	window.ShowAndRun()
}
