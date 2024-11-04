package main

import (
    "log"
    "math/rand"
    "time"

    "ball/src/scenes"
    "ball/src/views"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // Cargar la imagen del carro
    carImg, _, err := ebitenutil.NewImageFromFile("assets/images/carrote.png") // Asegúrate de que la ruta sea correcta
    if err != nil {
        log.Fatal(err)
    }

    // Inicializar el estado del juego
    game := scenes.NewGame()

    // Crear la interfaz de usuario, pasando el estado inicial y la imagen del carro
    gui := views.NewGameInterface(carImg)
    gui.Espacios = game.GetEspacios()
    gui.VehiculosRestantes = game.GetVehiculosRestantes()

    // Iniciar la simulación de vehículos
    go game.SimularLlegadaVehiculos(func(espacios [20]bool, vehiculosRestantes int, puertaOcupada bool, puertaSalidaOcupada bool) {
        gui.Espacios = espacios
        gui.VehiculosRestantes = vehiculosRestantes
        gui.PuertaOcupada = puertaOcupada // Actualizar el estado de la puerta de entrada
        gui.PuertaSalidaOcupada = puertaSalidaOcupada // Actualizar el estado de la puerta de salida
    })

    // Iniciar el juego
    if err := ebiten.RunGame(gui); err != nil {
        log.Fatal(err)
    }
}
