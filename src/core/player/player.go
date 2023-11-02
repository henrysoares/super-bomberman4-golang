package player

import (
	"fmt"
	"super-bomberman4/src/core/scenario"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CharacterTextureWidth  = 83
	CharacterTextureHeight = 121

	SingleCharacterWidth  = 22
	SingleCharacterHeight = 32

	PlayerFeetXOffset      = 20
	PlayerFeetYOffset      = 95
	PlayerFeetWidhtOffset  = 20
	PlayerFeetHeightOffset = 70
)

const (
	UPWARDS PlayerMovingDirection = iota
	BACKWARDS
	LEFTWARDS
	RIGHTWARDS
)

type PlayerMovingDirection uint8

type Player struct {
	PlayerNumber    int
	PlayerLifes     int
	PlayerFirePower int
	PlayerSpeed     float32
	PlayerScore     int

	time           int
	sprite         int
	texture        rl.Texture2D
	direction      int
	isPlayerMoving bool

	PosX float32
	PosY float32

	BombsSummoned int
}

func NewPlayer(number int) *Player {
	if number > 2 {
		panic("Não é possivel jogar com mais de 2 jogadores!")
	}

	return &Player{
		PlayerNumber:    number,
		PlayerLifes:     3,
		PlayerFirePower: 2,
		PlayerSpeed:     3.5,
		PlayerScore:     0,
		time:            0,
		sprite:          0,
		PosX:            0.0,
		PosY:            0.0,
		direction:       2,
		isPlayerMoving:  false,
		BombsSummoned:   0,
	}
}

// função que renderiza o jogador
func (p *Player) RenderPlayer() {

	textureRecDef := rl.NewRectangle(float32(SingleCharacterWidth*p.sprite), float32(SingleCharacterHeight*p.direction), 20, 25)
	playerRecOnGameDef := rl.NewRectangle(p.PosX, p.PosY, CharacterTextureWidth, CharacterTextureHeight)

	rl.DrawTexturePro(
		p.texture,
		textureRecDef,
		playerRecOnGameDef,
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)

	p.isPlayerMoving = false
}

// Set o spawn point do jogador
func (p *Player) SetSpawnPoint(x, y float32) {
	p.PosX = x
	p.PosY = y
}

// função que inicia o jogador
func (p *Player) InitPlayer() {

	switch p.PlayerNumber {
	case 1:
		p.texture = rl.LoadTexture("src/visual/characters/bomberman.png")
	default:
		p.texture = rl.LoadTexture("src/visual/characters/bomberman.png")
	}

	fmt.Printf("[Player] Player %v loaded successfully \n", p.PlayerNumber)
}

// função que mapeia todos inputs e keybinds possiveis do player
func (p *Player) HandlePlayerInputs() {
	if rl.IsKeyDown(rl.KeyW) {
		p.moveUpwards()
	}

	if rl.IsKeyDown(rl.KeyS) {
		p.moveBackwards()
	}

	if rl.IsKeyDown(rl.KeyA) {
		p.moveLeft()
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.moveRight()
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		p.BombsSummoned++
	}

	p.ackPlayerMotion()
}

// Retorna a hitbox dos pés do jogador
func (p Player) GetPlayerFeetHitbox() rl.Rectangle {

	return rl.NewRectangle(
		p.PosX+PlayerFeetXOffset,
		p.PosY+PlayerFeetYOffset,
		CharacterTextureWidth+PlayerFeetWidhtOffset,
		CharacterTextureHeight-PlayerFeetHeightOffset,
	)
}

func (p *Player) moveUpwards() {
	newCharacterPosY := p.PosY - p.PlayerSpeed
	p.moveCharacter(p.PosX, newCharacterPosY, 0)
}

func (p *Player) moveRight() {
	newCharacterPosX := p.PosX + p.PlayerSpeed
	p.moveCharacter(newCharacterPosX, p.PosY, 1)
}

func (p *Player) moveBackwards() {
	newCharacterPosY := p.PosY + p.PlayerSpeed
	p.moveCharacter(p.PosX, newCharacterPosY, 2)
}

func (p *Player) moveLeft() {
	newCharacterPosX := p.PosX - p.PlayerSpeed
	p.moveCharacter(newCharacterPosX, p.PosY, 3)
}

func (p *Player) moveCharacter(posX, posY float32, direction int) {
	newPlayerPosX := posX - p.PosX
	newPlayerPosY := posY - p.PosY

	if !p.WillPlayerColide(newPlayerPosX, newPlayerPosY) {
		p.PosX = float32(posX)
		p.PosY = float32(posY)
		p.isPlayerMoving = true
		p.direction = direction
	}
}

// Verifica se o usuario vai colidir com algum obstaculo.
func (p *Player) WillPlayerColide(newPlayerPosX, newPlayerPosY float32) bool {
	feet := p.GetPlayerFeetHitbox()

	x := feet.X
	y := feet.Y

	x += newPlayerPosX
	y += newPlayerPosY

	for _, border := range scenario.GetScenarioBoundaries() {
		if checkPlayerColision(x, y, border) {
			return true
		}
	}

	return false
}

// Verifica colisão do player
func checkPlayerColision(x, y float32, wall rl.Rectangle) bool {
	playerHitbox := rl.NewRectangle(x, y, 50, 20)

	return rl.CheckCollisionRecs(playerHitbox, wall)
}

// Faz a atualização dos parametros de tempo e sprite para atualizar o jogador.
func (p *Player) ackPlayerMotion() {
	if p.isPlayerMoving {
		if p.time%9 == 1 {
			p.sprite++
		}

		if p.sprite == 3 {
			p.sprite = 0
		}

		if p.time == 100 {
			p.time = 0
		}

		p.time++
	} else {
		p.sprite = 0
	}
}

func (p *Player) KillPlayer() {

}

// faz o descarregamento da textura do scenario
func (p Player) UnloadPlayer() {
	rl.UnloadTexture(p.texture)

	fmt.Printf("[Player] Player %v unloaded successfully \n", p.PlayerNumber)
}
