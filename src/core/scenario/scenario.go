package scenario

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TextureWidth  = 272
	TextureHeight = 208

	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 800

	QUADRANT_WIDHT_PROPORTION  = 4.70588235
	QUADRANT_HEIGHT_PROPORTION = 3.84615384

	SCNEARIO_QUADRANT_WIDTH  = 16 * QUADRANT_WIDHT_PROPORTION
	SCNEARIO_QUADRANT_HEIGHT = 16 * QUADRANT_HEIGHT_PROPORTION
)

var (
	arenaBackground rl.Texture2D
	textureRecDef   rl.Rectangle
	scenarioRecDef  rl.Rectangle

	initialScenarioSpawn = rl.NewVector2(0, 0)

	scenarioBoundaries = []rl.Rectangle{
		{X: 75.2, Y: 61.44, Width: 75, Height: 315},
		{X: 150.4, Y: 0, Width: 225.6, Height: 61.44},
		{X: 150.4, Y: 0, Width: 225.6, Height: 61.44},
		{X: 376, Y: 61.44, Width: 526.4, Height: 122.88},
		{X: 902.4, Y: 0, Width: 75.2 * 3, Height: 61.44},
		{X: 1128, Y: 61.44, Width: 75.2, Height: 61.44 * 6},
		{X: 150.4, Y: 368.64, Width: 75.2 * 2, Height: 61.44},
		{X: 376, Y: 368.64, Width: 75.2 * 7, Height: 61.44},
		{X: 977.6, Y: 368.64, Width: 75.2 * 2, Height: 61.44},
		{X: 150.4, Y: 430.08, Width: 75.2, Height: 61.44 * 5},
		{X: 225.6, Y: 737.28, Width: 75.2 * 11, Height: 61.44},
		{X: 526.4, Y: 430.08, Width: 75.2 * 3, Height: 61.44 * 2},
		{X: 526.4, Y: 614.4, Width: 75.2 * 3, Height: 61.44 * 2},
		{X: 1052.8, Y: 430.08, Width: 75.2, Height: 61.44 * 5},

		//scenario obstacles
		{X: 225.6, Y: 122.88, Width: 75.2, Height: 61.44},
		{X: 225.6, Y: 245.76, Width: 75.2 * 3, Height: 61.44},
		{X: 526.4, Y: 245.76, Width: 75.2, Height: 61.44},
		{X: 676.8, Y: 245.76, Width: 75.2, Height: 61.44},
		{X: 827.2, Y: 245.76, Width: 75.2, Height: 61.44},
		{X: 977.6, Y: 245.76, Width: 75.2, Height: 61.44},
		{X: 977.6, Y: 122.88, Width: 75.2, Height: 61.44},
		{X: 225.6, Y: 491.52, Width: 75.2, Height: 61.44},
		{X: 376, Y: 491.52, Width: 75.2, Height: 61.44},
		{X: 827, Y: 491.52, Width: 75.2, Height: 61.44},
		{X: 977.6, Y: 225.6, Width: 75.2, Height: 61.44},
		{X: 225.6, Y: 614.4, Width: 75.2, Height: 61.44},
		{X: 376, Y: 614.4, Width: 75.2, Height: 61.44},
		{X: 827.2, Y: 614.4, Width: 75.2, Height: 61.44},
		{X: 977.6, Y: 614.4, Width: 75.2, Height: 61.44},
	}
)

// realiza o carregamento inicial de texturas
func LoadScenario() {
	arenaBackground = rl.LoadTexture("src/visual/scenarious/mapa.png")
}

// inicializa variaveis do cenario
func InitScenario() {
	textureRecDef = rl.NewRectangle(0, 0, TextureWidth, TextureHeight)
	scenarioRecDef = rl.NewRectangle(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT)
}

// faz o calculo do quadrante pertencente de uma cordenada
func GetQuadrant(x, y float64) rl.Rectangle {
	xStartingPoint := math.Floor(x/SCNEARIO_QUADRANT_WIDTH) * SCNEARIO_QUADRANT_WIDTH
	yStartingPoint := math.Floor(y/SCNEARIO_QUADRANT_HEIGHT) * SCNEARIO_QUADRANT_HEIGHT

	return rl.NewRectangle(float32(xStartingPoint), float32(yStartingPoint), SCNEARIO_QUADRANT_WIDTH, SCNEARIO_QUADRANT_HEIGHT)
}

func GetScenarioBoundaries() []rl.Rectangle {
	return scenarioBoundaries
}

// faz o descarregamento da textura do scenario
func UnloadScenario() {
	rl.UnloadTexture(arenaBackground)
}

// renderiza o scenario na tela
func RenderScenario() {
	rl.DrawTexturePro(arenaBackground, textureRecDef, scenarioRecDef, initialScenarioSpawn, 0, rl.White)
}
