package main

import (
	//"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"log"
)

const (
	capacidadEstacionamiento = 20
	espacioAncho             = 80
	espacioAlto              = 40
	espacioMargen            = 10
	filas                     = 10 // Número total de filas
	columnas                  = 2  // Dos columnas
	totalVehiculos            = 100 // Número total de vehículos
)

type Game struct {
	espacios          [capacidadEstacionamiento]bool // true si ocupado, false si libre
	vehiculosRestantes int                             // Contador de vehículos restantes
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Dibujar fondo
	screen.Fill(colornames.Lightgray)

	// Calcular el desplazamiento horizontal para centrar el estacionamiento
	totalAncho := float64(columnas * (espacioAncho + espacioMargen) - espacioMargen)
	desplazamientoX := (640 - totalAncho) / 2 // Centrar en 640 píxeles de ancho

	// Dibujar el estacionamiento
	for fila := 0; fila < filas; fila++ {
		for columna := 0; columna < columnas; columna++ {
			indice := fila*columnas + columna
			x := desplazamientoX + float64(columna*(espacioAncho+espacioMargen))
			y := float64(fila * (espacioAlto + espacioMargen))

			// Dibuja el espacio
			drawEspacio(screen, x, y, g.espacios[indice])

			// Dibuja la división entre los espacios
			if columna < columnas-1 { // No dibujar línea después del último espacio en la fila
				drawDivision(screen, x+espacioAncho, y)
			}
		}
	}

	// Mostrar el contador de vehículos restantes
	drawContador(screen, g.vehiculosRestantes)
}

func drawEspacio(screen *ebiten.Image, x, y float64, ocupado bool) {
	espacio := ebiten.NewImage(espacioAncho, espacioAlto)
	if ocupado {
		espacio.Fill(colornames.Red) // Espacio ocupado
	} else {
		espacio.Fill(colornames.Green) // Espacio disponible
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(espacio, op)
}

func drawDivision(screen *ebiten.Image, x, y float64) {
	line := ebiten.NewImage(2, espacioAlto) // Línea divisoria
	line.Fill(colornames.Black)              // Color de la línea
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(line, op)
}

func drawContador(screen *ebiten.Image, vehiculosRestantes int) {
	text := "Vehículos restantes: " + strconv.Itoa(vehiculosRestantes)
	ebitenutil.DebugPrint(screen, text) // Usar DebugPrint directamente en el screen
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func (g *Game) simularLlegadaVehiculos() {
	g.vehiculosRestantes = totalVehiculos // Inicializar el contador
	for i := 0; i < totalVehiculos; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simular tiempo de llegada aleatorio

		// Buscar un espacio disponible
		for j := 0; j < capacidadEstacionamiento; j++ {
			if !g.espacios[j] { // Si el espacio está libre
				g.espacios[j] = true // Marcar como ocupado
				go g.vehiculoOcupando(j) // Simular el tiempo que el vehículo está en el estacionamiento
				break
			}
		}
		g.vehiculosRestantes-- // Disminuir el contador
	}
}

func (g *Game) vehiculoOcupando(indice int) {
	// Generar tiempo de permanencia aleatorio entre 5 y 10 segundos
	tiempoOcupacion := time.Duration(rand.Intn(6)+5) * time.Second
	time.Sleep(tiempoOcupacion) // Simular ocupación

	// Liberar el espacio
	g.espacios[indice] = false
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Inicializar generador de números aleatorios
	game := &Game{}
	go game.simularLlegadaVehiculos() // Comenzar la simulación de llegada de vehículos
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}