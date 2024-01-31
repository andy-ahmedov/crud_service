package rest

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CtxValue int

const (
	ctxUserID CtxValue = iota
)

func loggingMiddleware(c *gin.Context) {
	// log.Printf("%s: [%s] - %s ", time.Now().Format(time.RFC3339), r.Method, r.RequestURI)
	log.WithFields(log.Fields{
		"request": c.Request.Method,
		"uri":     c.Request.RequestURI,
	}).Info()
	c.Next()
}

func (h *Handler) authMiddleware(c *gin.Context) {
	token, err := getTokenFromRequest(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id, err := h.userService.ParseToken(c.Request.Context(), token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx := context.WithValue(c.Request.Context(), ctxUserID, id)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}

func getTokenFromRequest(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("empty auth header")
	}

	sub := strings.Split(token, " ")
	if len(sub) != 2 || sub[0] != "Beaver" {
		return "", errors.New("invalid auth header")
	}

	if len(sub[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return sub[1], nil
}
