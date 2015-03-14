package server

import (
	"math/rand"
	"net/http"
)

type Session struct {
	Id     string
	UserId int
}

// TODO: put this in redis or something
// global map is a quick hack
var sessionStore = map[string]*Session{}

const (
	SessionIDKey = "X-Askme-Token"
)

func NewSession() *Session {
	return &Session{
		Id:     newSessionId(),
		UserId: 0,
	}
}

func FetchSession(r *http.Request) *Session {
	key := r.Header.Get(SessionIDKey)
	if key == "" {
		key = newSessionId()
	}

	session := sessionStore[key]
	if session == nil {
		session = &Session{Id: key, UserId: 0}
		sessionStore[key] = session
	}

	return session
}

func SetSession(w http.ResponseWriter, s *Session) {
	sessionStore[s.Id] = s
	w.Header().Set(SessionIDKey, s.Id)
}

func DeleteSession(w http.ResponseWriter, s *Session) {
	sessionStore[s.Id] = nil
}

func newSessionId() string {
	return randomString(16)
}

func randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
