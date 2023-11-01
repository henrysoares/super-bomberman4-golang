package bomb

import (
	player "super-bomberman4/src/core/player"
	scenario "super-bomberman4/src/core/scenario"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BombWidthSize  = 16
	BombHeightSize = 16
)

type BombManager struct {
	Player                 *player.Player
	Bombs                  []*Bomb
	Explosions             []*Explosion
	MaximumSimultaneosBomb int
}

func NewBombManager(p *player.Player) *BombManager {
	InitBombs()
	return &BombManager{Player: p, MaximumSimultaneosBomb: 5}
}

// faz a adição da bomba no bomb manager para ser redenrizada assim que possivel
func (bm *BombManager) spawnBomb() {
	if bm.Player.BombsSummoned <= 0 {
		return
	}

	if len(bm.Bombs) < bm.MaximumSimultaneosBomb {
		bombSpawn := calculateBombSpawn(*bm.Player)

		bomb := NewBomb(bm.Player.PlayerFirePower, 0, bombSpawn)

		go bomb.CalculateExplosion()

		bm.Bombs = append(bm.Bombs, bomb)
	}

	bm.Player.BombsSummoned = 0
}

// renderiza bombas do scenario
func (bm *BombManager) renderBombs() {
	for _, bomb := range bm.Bombs {
		bm.drawSingleBomb(bomb)
	}

}

// faz a renderização de uma unica bomba
func (bm *BombManager) drawSingleBomb(bomb *Bomb) {
	if bomb.IsAlive {
		rl.DrawTexturePro(
			bomb.texture,
			rl.NewRectangle(float32(BombWidthSize*bomb.BombFrame)+float32(bomb.BombFrame), 0, BombWidthSize, BombHeightSize),
			rl.NewRectangle(float32(bomb.PosX), float32(bomb.PosY), scenario.SCNEARIO_QUADRANT_WIDTH, scenario.SCNEARIO_QUADRANT_HEIGHT),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)

		bomb.AckMotionTime()
	} else {
		bombIndex := bm.findIndexByBombID(bomb.BombID)
		bm.removeBomb(bombIndex)

		explosion := newExplosion()
		bm.Explosions = append(bm.Explosions, explosion)
	}

}

// remove bomba gerenciada pelo bomb manager
func (bm *BombManager) removeBomb(i int) {
	bombs := bm.Bombs
	bombs[i] = bombs[len(bombs)-1]
	bm.Bombs = bombs[:len(bombs)-1]
}

func (bm *BombManager) renderExplosions() {

}

func (bm *BombManager) ManageBombs() {
	bm.spawnBomb()
	bm.renderBombs()
	bm.renderExplosions()
}

func calculateBombSpawn(p player.Player) rl.Vector2 {
	playerFeet := p.GetPlayerFeetHitbox()
	quadrant := scenario.GetQuadrant(float64(playerFeet.X), float64(playerFeet.Y))

	return rl.NewVector2(quadrant.X, quadrant.Y)
}

// encontra bomba pelo ID da bomb
func (bm BombManager) findIndexByBombID(bombID string) int {
	for i, bomb := range bm.Bombs {
		if bomb.BombID == bombID {
			return i
		}
	}
	return -1
}
