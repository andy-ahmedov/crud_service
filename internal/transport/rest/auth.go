package rest

import (
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
// @Success 200 {string} gin.H "You are logged in!"
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

	user, err := h.userService.SignIn(c.Request.Context(), inp)
	if err != nil {
		if errors.Is(domain.ErrUserNotFound, err) {
			handlerErrUserNotFound("SignIn", err, c)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session, _ := h.store.Get(c.Request, "session-name")
	session.Values["user_id"] = user.ID
	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "You are logged in!"})
}

// @Summary Logout
// @Tags auth
// @Description User logout.
// @ID logout
// @Produce json
// @Success 200 {string} gin.H "You are logged out!"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /auth/logout [post]
func (h *Handler) logout(c *gin.Context) {
	session, err := h.store.Get(c.Request, "session-name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session.Options.MaxAge = -1

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "You are logged out!"})
}

func handlerErrUserNotFound(handler string, err error, c *gin.Context) {
	logError(handler, domain.ErrUserNotFound.Error(), err)
	c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrUserNotFound.Error()})
}
