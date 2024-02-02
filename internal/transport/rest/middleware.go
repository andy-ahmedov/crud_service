package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CtxValue int

const (
	ctxUserID CtxValue = iota
)

func loggingMiddleware(c *gin.Context) {
	log.WithFields(log.Fields{
		"request": c.Request.Method,
		"uri":     c.Request.RequestURI,
	}).Info()
	c.Next()
}

func (h *Handler) authMiddleware(c *gin.Context) {
	session, err := h.store.Get(c.Request, "session-name")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user_id, ok := session.Values["user_id"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	ctx := context.WithValue(c.Request.Context(), ctxUserID, user_id)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
