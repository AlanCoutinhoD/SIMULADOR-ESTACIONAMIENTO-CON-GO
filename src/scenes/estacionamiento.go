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
    Updates            chan GameUpdate // Canal para actualizaciones
}

type GameUpdate struct {
    Espacios           [CapacidadEstacionamiento]bool
    VehiculosRestantes int
    PuertaOcupada      bool
    PuertaSalidaOcupada bool
}

func NewGame() *Game {
    return &Game{
        vehiculosRestantes: 100,
        Updates:            make(chan GameUpdate), // Inicializar el canal
    }
}

func (g *Game) GetEspacios() [CapacidadEstacionamiento]bool {
    return g.espacios
}

func (g *Game) GetVehiculosRestantes() int {
    return g.vehiculosRestantes
}

// Función para simular la llegada de vehículos
func (g *Game) SimularLlegadaVehiculos() {
    for i := 0; i < 100; i++ {
        llegadaTiempo := time.Duration(rand.ExpFloat64() * 1000) * time.Millisecond
        time.Sleep(llegadaTiempo) // Simula llegada con distribución Poisson

        go g.llegarVehiculo()
    }
    g.vehiculosBloqueados.Wait() // Esperar a que todos los vehículos se vayan
}

// Lógica de llegada de un vehículo
func (g *Game) llegarVehiculo() {
    g.puertaEntrada.Lock()
    g.sendUpdate(true, false) // Actualizar el estado en el canal

    // Esperar 1 segundo en la entrada
    time.Sleep(1 * time.Second)

    // Buscar un espacio disponible
    espacioEncontrado := false
    for j := 0; j < CapacidadEstacionamiento; j++ {
        if !g.espacios[j] {
            g.espacios[j] = true // Marca el espacio como ocupado
            g.vehiculosBloqueados.Add(1)
            espacioEncontrado = true
            go g.ocuparEspacio(j) // Llama a ocuparEspacio
            break
        }
    }

    if espacioEncontrado {
        g.vehiculosRestantes-- // Disminuir la cantidad de vehículos restantes
    }
    g.puertaEntrada.Unlock()
}

// Función que simula el tiempo que un vehículo ocupa un espacio
func (g *Game) ocuparEspacio(indice int) {
    tiempoOcupacion := time.Duration(rand.Intn(3)+3) * time.Second // Tiempo aleatorio de ocupación entre 3 y 5 segundos
    time.Sleep(tiempoOcupacion)

    // Simular salida del vehículo
    g.puertaEntrada.Lock()
    g.espacios[indice] = false // Marca el espacio como libre
    g.puertaEntrada.Unlock()

    // Actualiza el estado de la puerta de salida
    g.sendUpdate(false, true)

    g.vehiculosBloqueados.Done() // Indicar que el vehículo ha salido
}

// Método para enviar actualizaciones a través del canal
func (g *Game) sendUpdate(puertaOcupada bool, puertaSalidaOcupada bool) {
    g.Updates <- GameUpdate{
        Espacios:           g.espacios,
        VehiculosRestantes: g.vehiculosRestantes,
        PuertaOcupada:      puertaOcupada,
        PuertaSalidaOcupada: puertaSalidaOcupada,
    }
}