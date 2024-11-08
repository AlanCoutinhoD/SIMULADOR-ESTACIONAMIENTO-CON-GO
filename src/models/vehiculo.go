// package models

// import (
//     "math/rand"
//     "sync"
//     "time"
//     "ball/src/scenes"
// )

// type Vehicle struct {
//     Index           int
//     Game            *scenes.Game
//     PuertaEntrada   *sync.Mutex
//     VehiculosBloqueados *sync.WaitGroup
// }

// func NewVehicle(game *scenes.Game, puertaEntrada *sync.Mutex, vehiculosBloqueados *sync.WaitGroup) *Vehicle {
//     return &Vehicle{
//         Game:            game,
//         PuertaEntrada:   puertaEntrada,
//         VehiculosBloqueados: vehiculosBloqueados,
//     }
// }

// // Función para la llegada de un vehículo
// func (v *Vehicle) LlegarVehiculo() {
//     v.PuertaEntrada.Lock()
//     v.Game.SendUpdate(true, false) // Actualizar el estado en el canal

//     // Esperar 1 segundo en la entrada
//     time.Sleep(1 * time.Second)

//     // Buscar un espacio disponible
//     espacioEncontrado := false
//     for j := 0; j < len(v.Game.GetEspacios()); j++ {
//         if !v.Game.GetEspacios()[j] {
//             v.Game.OcuparEspacio(j) // Ocupa el espacio en la estructura Game
//             v.VehiculosBloqueados.Add(1)
//             espacioEncontrado = true
//             go v.OcuparEspacio(j) // Llama a ocuparEspacio en el vehículo
//             break
//         }
//     }

//     if espacioEncontrado {
//         v.Game.DecrementarVehiculosRestantes()
//     }
//     v.PuertaEntrada.Unlock()
// }

// // Función que simula el tiempo que un vehículo ocupa un espacio
// func (v *Vehicle) OcuparEspacio(indice int) {
//     tiempoOcupacion := time.Duration(rand.Intn(3)+3) * time.Second // Tiempo aleatorio de ocupación entre 3 y 5 segundos
//     time.Sleep(tiempoOcupacion)

//     // Simular salida del vehículo
//     v.PuertaEntrada.Lock()
//     v.Game.LiberarEspacio(indice) // Marca el espacio como libre en la estructura Game
//     v.PuertaEntrada.Unlock()

//     // Actualiza el estado de la puerta de salida
//     v.Game.SendUpdate(false, true)

//     v.VehiculosBloqueados.Done() // Indicar que el vehículo ha salido
// }


//separar responsabilidades
//patron observador
//evitar funciones anonimas
