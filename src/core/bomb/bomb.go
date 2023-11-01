package bomb

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	uuid "github.com/google/uuid"
)

type Bomb struct {
	BombTick  int
	BombFrame int
	FirePower int
	BombType  int
	IsAlive   bool
	texture   rl.Texture2D
	BombID    string
	PosX      float32
	PosY      float32

	RenderExplosion bool
}

var (
	TextureFile rl.Texture2D
)

func NewBomb(firePower int, bombType int, cordinates rl.Vector2) *Bomb {

	return &Bomb{
		FirePower: firePower,
		BombType:  bombType,
		PosX:      cordinates.X,
		PosY:      cordinates.Y,
		texture:   TextureFile,
		IsAlive:   true,
		BombID:    uuid.NewString(),
	}
}

func InitBombs() {
	TextureFile = rl.LoadTexture("src/visual/bombs/bombs.png")

	fmt.Printf("[Bomb] Bomb loaded successfully \n")
}

func (b *Bomb) CalculateExplosion() {

}

func (b *Bomb) AckMotionTime() {

	if b.BombTick == 200 {
		b.IsAlive = false
		b.BombTick = 0
		b.RenderExplosion = true
		fmt.Println("BOOM")
	}

	if b.BombTick%12 == 1 {
		b.BombFrame++
	}

	if b.BombFrame == 3 {
		b.BombFrame = 0
	}

	b.BombTick++
}
