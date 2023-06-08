package lfg

import (
	"LFGbackend/graph/model"
	"sync"
	"time"
)

type SessionManager struct {
	sessions map[string]model.Player
	mutex    sync.Mutex
}

func NewManager() *SessionManager {
	m := &SessionManager{
		sessions: make(map[string]model.Player),
		mutex:    sync.Mutex{},
	}
	m.Set("dev", model.Player{})
	return m
}

func (sm *SessionManager) Get(clientId string) (model.Player, bool) {
	defer sm.mutex.Unlock()
	sm.mutex.Lock()
	data, ok := sm.sessions[clientId]
	return data, ok
}

func (sm *SessionManager) Set(clientId string, pd model.Player) {

	sm.mutex.Lock()
	sm.sessions[clientId] = pd
	sm.mutex.Unlock()

	go func() {
		time.Sleep(time.Hour * 24)
		sm.mutex.Lock()
		delete(sm.sessions, clientId)
		sm.mutex.Unlock()
	}()
}
