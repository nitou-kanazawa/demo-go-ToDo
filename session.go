package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"time"
)

const cookiesSeddionId = "sessionId"

// セッションが開始されていることを保証する
func ensureSession(w http.ResponseWriter, r *http.Request) (string, error){
	c, err := r.Cookie(cookiesSeddionId)
	
	// CookieにセッションIDが入ってない場合は，新しく発行して返す
	if err == http.ErrNoCookie{
		sessionId, err := startSession(w)
		return sessionId, err
	}

	// CokkieにセッションIDが入っている場合は，それを返す
	if err == nil {
		 sessionId := c.Value
		 return sessionId, nil
	}

	return "", nil
}

// セッションを開始する
func startSession(w http.ResponseWriter) (string, error){
	sessionId, err := makeSessionId()
	if err != nil{
		return "",err
	}

	cookie := &http.Cookie{
		Name: cookiesSeddionId,
		Value: sessionId,
		Expires: time.Now().Add(1800 * time.Second),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	return sessionId, nil
}

// セッションIDを生成する
func makeSessionId() (string, error){
	randByte := make([]byte, 16)
	if _,err := io.ReadFull(rand.Reader, randByte); err != nil {
			return "", err
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(randByte), nil
}