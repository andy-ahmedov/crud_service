package rest

import (
	"context"
	"errors"
	"net/http"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description Registration of a new user in the system.
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body domain.SignUpInput true "User info"
// @Success 200 {string} gin.H "The user has been successfully registered."
// @Failure 400 {object} errResponse "Bad Request"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var user domain.SignUpInput

	if err := c.ShouldBindJSON(&user); err != nil {
		logError("signUp", "Invalid format", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userService.SignUp(c.Request.Context(), user); err != nil {
		logError("signUp", "Internal Service Error", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// @Summary SignIn
// @Tags auth
// @Description User authentication by email and password.
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body domain.SignInInput true "User info"
// @Success 200 {string} gin.H "The JWT token was successfully generated."
// @Failure 400 {object} errResponse "Bad Request"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var inp domain.SignInInput

	if err := c.ShouldBindJSON(&inp); err != nil {
		logError("signIn", "Invalid format", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.SignIn(context.TODO(), inp)
	if err != nil {
		if errors.Is(domain.ErrUserNotFound, err) {
			handlerErrUserNotFound("SignIn", err, c)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func handlerErrUserNotFound(handler string, err error, c *gin.Context) {
	logError(handler, domain.ErrUserNotFound.Error(), err)
	c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrUserNotFound.Error()})
}