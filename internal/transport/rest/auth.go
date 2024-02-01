package rest

import (
	"errors"
	"net/http"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	accessToken, refreshToken, err := h.userService.SignIn(c.Request.Context(), inp)
	if err != nil {
		if errors.Is(domain.ErrUserNotFound, err) {
			handlerErrUserNotFound("SignIn", err, c)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.Request.Header.Add("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))
	c.SetCookie("refresh-token", refreshToken, 0, "/auth", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": accessToken})
}

// @Summary Refresh
// @Tags auth
// @Description Refresh token update.
// @ID refresh
// @Produce json
// @Success 200 {string} gin.H "Refresh token has been successfully updated."
// @Failure 400 {object} errResponse "Bad Request"
// @Failure 500 {object} errResponse "Internal Server Error"
// @Router /auth/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	cookie, err := c.Request.Cookie("refresh-token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("%v", cookie.Value)

	accesToken, refreshToken, err := h.userService.RefreshTokens(c.Request.Context(), cookie.Value)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// c.Request.Header.Add("Set-Cookie", fmt.Sprintf("refresh-token='%s'; HttpOnly", refreshToken))
	c.SetCookie("refresh-token", refreshToken, 0, "/auth", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": accesToken})
}

func handlerErrUserNotFound(handler string, err error, c *gin.Context) {
	logError(handler, domain.ErrUserNotFound.Error(), err)
	c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrUserNotFound.Error()})
}
