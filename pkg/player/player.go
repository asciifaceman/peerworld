package player

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/asciifaceman/peerworld/pkg/namegen"
)

func NewPlayer(leader bool) *Player {
	now := time.Now()

	name := "GameLeader"
	if !leader {
		sex := rand.Intn(2)
		if sex == 0 {
			name = namegen.RandomFemmeName()
		} else {
			name = namegen.RandomMascName()
		}
	}

	p := &Player{
		Name:   name,
		born:   now,
		maxAge: 10,
	}
	return p
}

type Player struct {
	Name   string
	born   time.Time
	maxAge time.Duration
}

func (p *Player) FmtFullName() string {
	return fmt.Sprintf("%s", p.Name)
}
