package api

import (
	"crm-system/pkg/logger"
	"crm-system/pkg/model"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"net/http"
	"strings"
)

type UserHandler struct {
	api *api
}

func NewUserHandler(a *api) *UserHandler {
	return &UserHandler{
		api: a,
	}
}

// UpdateInfo
// @Summary update user info
// @Produce json
// @Tags User
// @Security ApiKeyAuth
// @Param User  body model.User  true "User"
// @Success 200 {object} model.User
// @Failure 400 {object} errors.UIResponseErrorBadRequest
// @Router /api/v1/user/update-info [patch]
//
//nolint:varnamelen
func (h *UserHandler) UpdateInfo(c *gin.Context) {
	user := &model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		logger.Errorf("UpdatePersonalInfo.ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, model.ErrInvalidBody)

		return
	}

	user.UserID, err = h.getUserIDFromHeader(c)
	if err != nil {
		logger.Errorf("UpdatePersonalInfo.getUserRoleFromHeader", err)
		c.JSON(http.StatusUnauthorized, model.ErrUnauthorized)

		return
	}

	err = h.api.postgresStore.User.Update(user)
	if err != nil {
		logger.Errorf("UpdatePersonalInfo.Update", err)
		c.JSON(http.StatusInternalServerError, model.ErrUnhealthy)

		return
	}

	c.JSON(http.StatusOK, user)
}

// Get
// @Summary get user info
// @Produce json
// @Tags User
// @Security ApiKeyAuth
// @Success 200 {object} model.User
// @Failure 400 {object} errors.UIResponseErrorBadRequest
// @Router /api/v1/user [get]
//
//nolint:varnamelen
func (h *UserHandler) Get(c *gin.Context) {
	userID, err := h.getUserIDFromHeader(c)
	if err != nil {
		logger.Errorf("Get.getUserRoleFromHeader", err)
		c.JSON(http.StatusUnauthorized, model.ErrUnauthorized)

		return
	}

	user, err := h.api.postgresStore.User.Get(userID)
	if err != nil {
		logger.Errorf("UpdatePersonalInfo. uuid.FromString", err)

		c.JSON(http.StatusInternalServerError, model.ErrUnhealthy)

		return
	}

	c.JSON(http.StatusOK, user)
}

//nolint:gocritic
func (h *UserHandler) getUserIDFromHeader(c *gin.Context) (uuid.UUID, error) {
	token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)
	userID, err := h.api.auth.GetUserID(token)

	return userID, err
}
