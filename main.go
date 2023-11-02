package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	bomb "super-bomberman4/src/core/bomb"
	player "super-bomberman4/src/core/player"
	scenario "super-bomberman4/src/core/scenario"
)

const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 800
	GAME_NAME     = "Super bomberman 4"
	TARGET_FPS    = 60
)

var (
	player1     *player.Player
	bombManager *bomb.BombManager
)

// inicia todos os recursos necessarios do game.
func initGame() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, GAME_NAME)
	rl.SetTargetFPS(TARGET_FPS)

	scenario.LoadScenario()
	scenario.InitScenario()

	player1 = player.NewPlayer(1)
	player1.InitPlayer()

	player1.SetSpawnPoint(150, 0)

	bombManager = bomb.NewBombManager(player1)
}

// faz o update do jogo
func updateGame() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	scenario.RenderScenario()

	player1.HandlePlayerInputs()
	bombManager.ManageBombs()
	player1.RenderPlayer()

	rl.EndDrawing()
}

// inicia o processo de desligamento do game
func closeGame() {
	scenario.UnloadScenario()
	player1.UnloadPlayer()

	rl.CloseWindow()
}

func main() {
	initGame()

	for !rl.WindowShouldClose() {
		updateGame()
	}

	closeGame()
}
