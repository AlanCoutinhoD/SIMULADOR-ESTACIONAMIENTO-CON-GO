package views

import (
    "strconv"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "golang.org/x/image/colornames"
)

const (
    EspacioAncho  = 80
    EspacioAlto   = 40
    EspacioMargen = 10
    Filas         = 10
    Columnas      = 2
)

type GameInterface struct {
    Espacios          [20]bool
    VehiculosRestantes int
}

func (gi *GameInterface) Update() error {
    return nil
}

func (gi *GameInterface) Draw(screen *ebiten.Image) {
    screen.Fill(colornames.Lightgray)
    totalAncho := float64(Columnas*(EspacioAncho+EspacioMargen) - EspacioMargen)
    desplazamientoX := (640 - totalAncho) / 2

    for fila := 0; fila < Filas; fila++ {
        for columna := 0; columna < Columnas; columna++ {
            indice := fila*Columnas + columna
            x := desplazamientoX + float64(columna*(EspacioAncho+EspacioMargen))
            y := float64(fila * (EspacioAlto + EspacioMargen))
            drawEspacio(screen, x, y, gi.Espacios[indice])

            if columna < Columnas-1 {
                drawDivision(screen, x+EspacioAncho, y)
            }
        }
    }
    drawContador(screen, gi.VehiculosRestantes)
}

func (gi *GameInterface) Layout(outsideWidth, outsideHeight int) (int, int) {
    return 640, 400
}

func drawEspacio(screen *ebiten.Image, x, y float64, ocupado bool) {
    espacio := ebiten.NewImage(EspacioAncho, EspacioAlto)
    if ocupado {
        espacio.Fill(colornames.Red)
    } else {
        espacio.Fill(colornames.Green)
    }
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(x, y)
    screen.DrawImage(espacio, op)
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
    ebitenutil.DebugPrint(screen, text)
}
