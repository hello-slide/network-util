package networkutil

import (
	"net/http"
	"net/url"
	"strings"
)

type TokenOperation struct {
	CookieOp CookieOperation
}

// Create tokenOp.
//
// Arguments;
//	apiUrl {string} - URL of the API you are using.
func NewTokenOp(apiUrl string) (*TokenOperation, error) {
	u, err := url.Parse(apiUrl)
	if err != nil {
		return nil, err
	}
	splittedHost := strings.Split(u.String(), ".")
	hostLen := len(splittedHost)
	domain := strings.Join(splittedHost[hostLen-2:], ".")

	cookieOp := NewCookieOp(domain)
	return &TokenOperation{
		CookieOp: *cookieOp,
	}, nil
}

// Set the refresh token.
//
// Arguments:
//	w {http.ResponseWriter} - http writer.
//	refreshToken {string} - refresh token.
func (t *TokenOperation) SetRefreshToken(w http.ResponseWriter, refreshToken string) {
	t.CookieOp.Set(w, "refresh_token", refreshToken, 24*30)
}

// Set the session token.
//
// Arguments:
//	w {http.ResponseWriter} - http writer.
//	sessionToken {string} - session token.
func (t *TokenOperation) SetSessionToken(w http.ResponseWriter, sessionToken string) {
	t.CookieOp.Set(w, "session_token", sessionToken, 6)
}

// Get the refresh token.
//
// Arguments:
//	r {*http.Request} - http requests.
//
// Returns:
//	{string} - refresh token.
func (t *TokenOperation) GetRefreshToken(r *http.Request) (string, error) {
	return t.CookieOp.Get(r, "refresh_token")
}

// Get the session token.
//
// Arguments:
//	r {*http.Request} - http requests.
//
// Returns:
//	{string} - session token.
func (t *TokenOperation) GetSessionToken(r *http.Request) (string, error) {
	return t.CookieOp.Get(r, "session_token")
}

// Delete session and refresh tokens.
//
// Arguments:
//	w {http.ResponseWriter} - http writer.
//	r {*http.Request} - http requests.
func (t *TokenOperation) DeleteToken(w http.ResponseWriter, r *http.Request) error {
	if err := t.CookieOp.Delete(w, r, "session_token"); err != nil {
		return err
	}

	if err := t.CookieOp.Delete(w, r, "refresh_token"); err != nil {
		return err
	}
	return nil
}
