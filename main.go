package main

import (
    "log"
    "math/rand"
    "time"

    "ball/src/scenes"
    "ball/src/views"

    "github.com/hajimehoshi/ebiten/v2"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // Inicializar el estado del juego
    game := scenes.NewGame()

    // Crear la interfaz de usuario, pasando el estado inicial
    gui := &views.GameInterface{
        Espacios:          game.GetEspacios(),
        VehiculosRestantes: game.GetVehiculosRestantes(),
    }

    // Iniciar la simulación de vehículos
    go game.SimularLlegadaVehiculos(func(espacios [20]bool, vehiculosRestantes int) {
        gui.Espacios = espacios
        gui.VehiculosRestantes = vehiculosRestantes
    })

    if err := ebiten.RunGame(gui); err != nil {
        log.Fatal(err)
    }
}
