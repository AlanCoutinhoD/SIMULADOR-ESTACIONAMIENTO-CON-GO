package models

// Vehiculo representa un vehículo que intenta entrar y salir del estacionamiento.
type Vehiculo struct {
    ID int
}

func NewVehiculo(id int) *Vehiculo {
    return &Vehiculo{ID: id}
}


//separar responsabilidades
//patron observador
//evitar funciones anonimas