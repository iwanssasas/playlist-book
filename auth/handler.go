package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"PLAYLISTBOOK/utils"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return Handler{
		service: s,
	}
}

func (h Handler) handleRegister(c *gin.Context) {
	ctx := context.Background()
	var params RegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	data, err := h.service.Register(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}

func (h Handler) handleLogin(c *gin.Context) {
	ctx := context.Background()
	var params LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	data, err := h.service.Login(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}

func (h Handler) handlePing(c *gin.Context) {
	ctx := context.Background()
	data, err := h.service.Ping(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}

func (h Handler) handleAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Response("Admin Only!"))
}

func (h Handler) handleOauthGoogle(c *gin.Context) {
	url := getGoogleOauth2().AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h Handler) handleGoogleCallback(c *gin.Context) {
	ctx := context.Background()
	content, err := h.service.googleCallback(ctx, c.Query("state"), c.Query("code"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	data, err := h.service.registerGoogleAuth(ctx, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}

func (h Handler) handleGetProfile(c *gin.Context) {
	ctx := context.Background()
	id := c.GetString("id")

	data, err := h.service.GetUSer(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))

}

func (h Handler) handleEditProfile(c *gin.Context) {
	var params UpdateProfileParam
	ctx := context.Background()
	id := c.GetString("id")

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	data, err := h.service.EditUser(ctx, params, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))

}

func (h Handler) handleDeleteProfile(c *gin.Context) {
	id := c.GetString("id")
	err := h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response("YourAccountHasDeleted"))
}
