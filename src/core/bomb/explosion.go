package bomb

import rl "github.com/gen2brain/raylib-go/raylib"

type Explosion struct {
	ExplosionsTailUpper int
	ExplosionsTailDown  int
	ExplosionsTailRight int
	ExplosionsTailLeft  int
	Center              rl.Vector2
	ExplosionsUpper     int
	ExplosionsDown      int
	ExplosionsRight     int
	ExplosionsLeft      int
}

func newExplosion() *Explosion {
	return &Explosion{}
}
