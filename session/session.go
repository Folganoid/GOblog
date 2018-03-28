package session

import (
	"../utils"
	"github.com/codegangsta/martini"
	"net/http"
	"time"
)

const (
	COOKIE_NAME = "sessionId"
)

type Session struct {
	id       string
	Username string
}

type SessionStore struct {
	data map[string]*Session
}

func NewSessionStore() *SessionStore {
	s := new(SessionStore)
	s.data = make(map[string]*Session)

	return s
}

func (store *SessionStore) Get(sessionId string) *Session {
}

func (store *SessionStore) Set(session *Session) {
}

func ensureCookie(r *http.Request, w http.ResponseWriter) string {
	cookie := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		return cookie.Value
	}
	sessionId := utils.GenerateId()
	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}
	http.SetCookie(w, cookie)

	return sessionId
}

var sessionStore = NewSessionStore()

func Middleware(ctx martini.Context, r *http.Request, w http.ResponseWriter) {
	sessionId := ensureCookie(r, w)
	ctx.Next()
}
