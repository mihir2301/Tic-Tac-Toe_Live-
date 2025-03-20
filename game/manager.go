package game

import (
	"errors"
	"sync"
)

type Manager struct {
	mu            sync.Mutex
	Games         map[string]*Game
	Activeplayers map[string]string
}

func NewManager() *Manager {
	return &Manager{
		Games:         make(map[string]*Game),
		Activeplayers: make(map[string]string),
	}
}

//create a new game seesion

func (m *Manager) CreateGame(gameID string, playerID string) (*Game, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exist := m.Activeplayers[playerID]
	if exist {
		return nil, errors.New("palyer already exist in active game")
	}
	game := NewGame(gameID)
	err := game.AddPlayer(playerID)
	if err != nil {
		return nil, err
	}
	m.Games[gameID] = game
	m.Activeplayers[playerID] = gameID

	return game, nil
}

// End the Game and remove palyers from active status
func (m *Manager) EndGame(gameID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	game, exist := m.Games[gameID]
	if !exist {
		return
	}

	for _, player := range game.Players {
		delete(m.Activeplayers, player)
	}

	delete(m.Games, gameID)
}
