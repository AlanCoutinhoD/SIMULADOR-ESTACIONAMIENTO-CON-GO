package scenes

import (
    "math/rand"
    "sync"
    "time"
)

const CapacidadEstacionamiento = 20

type Game struct {
    espacios           [CapacidadEstacionamiento]bool
    vehiculosRestantes int
    puertaEntrada      sync.Mutex
    vehiculosBloqueados sync.WaitGroup
}

func NewGame() *Game {
    return &Game{
        vehiculosRestantes: 100,
    }
}

func (g *Game) GetEspacios() [CapacidadEstacionamiento]bool {
    return g.espacios
}

func (g *Game) GetVehiculosRestantes() int {
    return g.vehiculosRestantes
}

func (g *Game) SimularLlegadaVehiculos(actualizarUI func([CapacidadEstacionamiento]bool, int)) {
    for i := 0; i < 100; i++ {
        llegadaTiempo := time.Duration(rand.ExpFloat64() * 1000) * time.Millisecond
        time.Sleep(llegadaTiempo) // Simula llegada con distribución Poisson

        go func() {
            for {
                g.puertaEntrada.Lock()

                // Buscar un espacio disponible
                espacioEncontrado := false
                for j := 0; j < CapacidadEstacionamiento; j++ {
                    if !g.espacios[j] {
                        g.espacios[j] = true
                        g.vehiculosBloqueados.Add(1)
                        espacioEncontrado = true
                        go g.ocuparEspacio(j, actualizarUI)
                        break
                    }
                }

                if espacioEncontrado {
                    g.vehiculosRestantes--
                    actualizarUI(g.espacios, g.vehiculosRestantes)
                    g.puertaEntrada.Unlock()
                    break // Salir del ciclo si el vehículo consiguió un espacio
                } else {
                    g.puertaEntrada.Unlock()
                    time.Sleep(500 * time.Millisecond) // Espera un tiempo antes de reintentar
                }
            }
        }()
    }
    g.vehiculosBloqueados.Wait() // Esperar a que todos los vehículos se vayan
}

func (g *Game) ocuparEspacio(indice int, actualizarUI func([CapacidadEstacionamiento]bool, int)) {
    tiempoOcupacion := time.Duration(rand.Intn(3)+3) * time.Second
    time.Sleep(tiempoOcupacion)

    g.puertaEntrada.Lock()
    g.espacios[indice] = false
    actualizarUI(g.espacios, g.vehiculosRestantes)
    g.puertaEntrada.Unlock()
    g.vehiculosBloqueados.Done() // Indicar que el vehículo ha salido
}
