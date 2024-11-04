package views

import (
    //"log"
    "strconv"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "golang.org/x/image/colornames"
)

const (
    EspacioAncho    = 60
    EspacioAlto     = 30
    EspacioMargen   = 8
    Filas           = 10
    Columnas        = 2
    AnchoPantalla   = 640
    AltoPantalla    = 500
    AnchoEntrada     = 100
    AltoEntrada      = 30
    AnchoSalida      = 100
    AltoSalida       = 30
    OffsetEntradaX   = 250 // Nueva constante para mover la entrada
    OffsetSalidaX    = 400 // Nueva constante para mover la salida
)

type GameInterface struct {
    Espacios           [20]bool
    VehiculosRestantes int
    PuertaOcupada      bool
    PuertaSalidaOcupada bool // Añadido para el estado de la puerta de salida
    CarImage           *ebiten.Image // Imagen del carro
}

// Nueva función para cargar la interfaz del juego, aceptando la imagen del carro
func NewGameInterface(carImg *ebiten.Image) *GameInterface {
    return &GameInterface{
        CarImage: carImg,
    }
}

func (gi *GameInterface) Update() error {
    return nil
}

func (gi *GameInterface) Draw(screen *ebiten.Image) {
    screen.Fill(colornames.Lightgray)
    totalAncho := float64(Columnas*(EspacioAncho+EspacioMargen) - EspacioMargen)
    desplazamientoX := (AnchoPantalla - totalAncho) / 2

    // Dibujar la puerta de entrada
    drawEntrada(screen, gi.PuertaOcupada)

    // Dibujar los espacios de estacionamiento
    for fila := 0; fila < Filas; fila++ {
        for columna := 0; columna < Columnas; columna++ {
            indice := fila*Columnas + columna
            if indice >= len(gi.Espacios) {
                continue // Asegurarse de no exceder el límite del array
            }
            x := desplazamientoX + float64(columna*(EspacioAncho+EspacioMargen))
            y := float64(fila*(EspacioAlto+EspacioMargen)) + AltoEntrada + 20
            drawEspacio(screen, x, y, gi.Espacios[indice], gi.CarImage)

            if columna < Columnas-1 {
                drawDivision(screen, x+EspacioAncho, y)
            }
        }
    }

    // Dibujar la puerta de salida
    drawSalida(screen, gi.PuertaSalidaOcupada)

    // Dibujar el contador de vehículos restantes
    drawContador(screen, gi.VehiculosRestantes)
}

func (gi *GameInterface) Layout(outsideWidth, outsideHeight int) (int, int) {
    return AnchoPantalla, AltoPantalla
}

func drawEntrada(screen *ebiten.Image, ocupada bool) {
    entrada := ebiten.NewImage(AnchoEntrada, AltoEntrada)
    if ocupada {
        entrada.Fill(colornames.Red) // Puerta ocupada
    } else {
        entrada.Fill(colornames.Green) // Puerta libre
    }
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(OffsetEntradaX, 10) // Ubicar en la parte superior y a la derecha
    screen.DrawImage(entrada, op)
    ebitenutil.DebugPrintAt(screen, "Entrada", OffsetEntradaX+15, 15)
}

func drawSalida(screen *ebiten.Image, ocupada bool) {
    salida := ebiten.NewImage(AnchoSalida, AltoSalida)
    if ocupada {
        salida.Fill(colornames.Blue) // Puerta ocupada
    } else {
        salida.Fill(colornames.Green) // Puerta libre
    }
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(OffsetSalidaX, 10) // Ubicar en la parte superior y a la derecha
    screen.DrawImage(salida, op)
    ebitenutil.DebugPrintAt(screen, "Salida", OffsetSalidaX+15, 15)
}

func drawEspacio(screen *ebiten.Image, x, y float64, ocupado bool, carImage *ebiten.Image) {
    if ocupado {
        // Dibujar el carro en el espacio ocupado
        op := &ebiten.DrawImageOptions{}
        op.GeoM.Scale(float64(EspacioAncho)/float64(carImage.Bounds().Dx()), float64(EspacioAlto)/float64(carImage.Bounds().Dy()))
        op.GeoM.Translate(x, y)
        screen.DrawImage(carImage, op)
    } else {
        espacio := ebiten.NewImage(EspacioAncho, EspacioAlto)
        espacio.Fill(colornames.Green)
        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(x, y)
        screen.DrawImage(espacio, op)
    }
}

func drawDivision(screen *ebiten.Image, x, y float64) {
    line := ebiten.NewImage(2, EspacioAlto)
    line.Fill(colornames.Black)
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(x, y)
    screen.DrawImage(line, op)
}

func drawContador(screen *ebiten.Image, vehiculosRestantes int) {
    text := "Vehículos restantes: " + strconv.Itoa(vehiculosRestantes)
    ebitenutil.DebugPrintAt(screen, text, 10, AltoPantalla-20)
}
