package session

import (
	"github.com/kscarlett/muzsh/game"
	"github.com/kscarlett/muzsh/player"
)

var (
	Data *Session
)

type Session struct {
	Player *player.Player
	World  *game.World
}

func New() *Session {
	return &Session{
		Player: player.New(),
		World:  game.NewWorld(),
	}
}
