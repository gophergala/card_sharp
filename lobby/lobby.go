package lobby

import (
	"fmt"
	"math/rand"
	"time"

	"git.andrewcsellers.com/acsellers/card_sharp/store"
)

var activeLobbies map[string]*Lobby

func init() {
	activeLobbies = map[string]*Lobby{}
}

type Lobby struct {
	ID       string
	Name     string
	Players  []string
	Creation time.Time
	Game     *store.Deck
}

func Create(d store.Deck) *Lobby {
	l := Lobby{
		ID:       fmt.Sprint(rand.Intn(100000)),
		Name:     "Lobby Name Here",
		Creation: time.Now(),
		Game:     &d,
	}

	// make sure we aren't reusing another lobby's name
	_, ok := activeLobbies[l.ID]
	for ok {
		l.ID = fmt.Sprint(rand.Intn(100000))
		_, ok = activeLobbies[l.ID]
	}

	activeLobbies[l.ID] = &l
	return &l
}

func Find(id string) *Lobby {
	return activeLobbies[id]
}
