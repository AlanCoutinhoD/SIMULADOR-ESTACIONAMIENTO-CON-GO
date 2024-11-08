package main

import (
    "log"
    //"math/rand"
   // "time"

    "ball/src/scenes"
    "ball/src/views"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
   // rand.Seed(time.Now().UnixNano())

    
    carImg, _, err := ebitenutil.NewImageFromFile("assets/images/carrote.png") 
    if err != nil {
        log.Fatal(err)
    }

    // Inicializar el estado del simulador
    game := scenes.NewGame()

    // Crear la interfaz de usuario, pasando el estado inicial y la imagen del carro
    gui := views.NewGameInterface(carImg)

    // Iniciar el goroutine para escuchar actualizaciones
    go gui.ListenToUpdates(game.Updates)

    // Iniciar la simulación de vehículos
    go game.SimularLlegadaVehiculos()

    // Iniciar 
    if err := ebiten.RunGame(gui); err != nil {
        log.Fatal(err)
    }
}