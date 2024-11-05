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

// Función para simular la llegada de vehículos
func (g *Game) SimularLlegadaVehiculos(actualizarUI func([CapacidadEstacionamiento]bool, int, bool, bool)) {
    for i := 0; i < 100; i++ {
        llegadaTiempo := time.Duration(rand.ExpFloat64() * 1000) * time.Millisecond
        time.Sleep(llegadaTiempo) // Simula llegada con distribución Poisson

        go g.llegarVehiculo(actualizarUI)
    }
    g.vehiculosBloqueados.Wait() // Esperar a que todos los vehículos se vayan
}

// Lógica de llegada de un vehículo
func (g *Game) llegarVehiculo(actualizarUI func([CapacidadEstacionamiento]bool, int, bool, bool)) {
    g.puertaEntrada.Lock()
    actualizarUI(g.espacios, g.vehiculosRestantes, true, false) // Indicar que la puerta de entrada está ocupada

    // Esperar 1 segundo en la entrada
    time.Sleep(1 * time.Second)

    // Buscar un espacio disponible
    espacioEncontrado := false
    for j := 0; j < CapacidadEstacionamiento; j++ {
        if !g.espacios[j] {
            g.espacios[j] = true // Marca el espacio como ocupado
            g.vehiculosBloqueados.Add(1)
            espacioEncontrado = true
            go g.ocuparEspacio(j, actualizarUI) // Llama a ocuparEspacio
            break
        }
    }

    if espacioEncontrado {
        g.vehiculosRestantes-- // Disminuir la cantidad de vehículos restantes
    }
    g.puertaEntrada.Unlock()
}

// Función que simula el tiempo que un vehículo ocupa un espacio
func (g *Game) ocuparEspacio(indice int, actualizarUI func([CapacidadEstacionamiento]bool, int, bool, bool)) {
    tiempoOcupacion := time.Duration(rand.Intn(3)+3) * time.Second // Tiempo aleatorio de ocupación entre 3 y 5 segundos
    time.Sleep(tiempoOcupacion)

    // Simular salida del vehículo
    g.puertaEntrada.Lock()
    g.espacios[indice] = false // Marca el espacio como libre
    g.puertaEntrada.Unlock()

    // Actualiza el estado de la puerta de salida
    actualizarUI(g.espacios, g.vehiculosRestantes, false, true)

    g.vehiculosBloqueados.Done() // Indicar que el vehículo ha salido
}
