// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"strconv"

	"github.com/chinmobi/gin-mvc/restful"
	"github.com/chinmobi/gin-mvc/service"
	"github.com/chinmobi/gin-mvc/service/dto"

	"github.com/gin-gonic/gin"
)

// NOTE: The user controller is just for demo, could to be removed for real project.

type UserController struct {
	services  service.Supplier
	userSvc  *service.UserService
}

func NewUserController(services service.Supplier) *UserController {
	ctrl := &UserController{
		services: services,
		userSvc:  services.GetUserService(),
	}
	return ctrl
}

// --- Handler methods ---

// POST /users
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var userDTO dto.UserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		restful.RespBadRequest(c, err)
		return
	}

	user, err := ctrl.getUserSvc().CreateUser(&userDTO)
	if err != nil {
		restful.RespServiceError(c, err)
		return
	}

	restful.RespCreatedDataEntity(c, user)
}

// GET /users
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.getUserSvc().FindAllUsers()
	if err != nil {
		restful.RespServiceError(c, err)
		return
	}

	restful.RespDataEntity(c, users)
}

// GET /users/:uid
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	uid, err := getUserIdParam(c)
	if err != nil {
		restful.RespBadRequest(c, err)
		return
	}

	user, err := ctrl.getUserSvc().FindUserByID(uid)
	if err != nil {
		restful.RespServiceError(c, err)
		return
	}

	restful.RespDataEntity(c, user)
}

// PATCH /users/:uid
// PUT /users/:uid
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	uid, err := getUserIdParam(c)
	if err != nil {
		restful.RespBadRequest(c, err)
		return
	}

	var userDTO dto.UserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		restful.RespBadRequest(c, err)
		return
	}

	user, err := ctrl.getUserSvc().UpdateUser(uid, &userDTO)
	if err != nil {
		restful.RespServiceError(c, err)
		return
	}

	restful.RespDataEntity(c, user)
}

// DELETE /users/:uid
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	uid, err := getUserIdParam(c)
	if err != nil {
		restful.RespBadRequest(c, err)
		return
	}

	_, err = ctrl.getUserSvc().DeleteUser(uid)
	if err != nil {
		restful.RespServiceError(c, err)
		return
	}

	restful.RespNoContent(c)
}

// --- Helper methods ---

func (ctrl *UserController) getUserSvc() *service.UserService {
	return ctrl.userSvc
}

func parseUserId(id string) (uint32, error) {
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(uid), nil
}

func getUserIdParam(c *gin.Context) (uint32, error) {
	return parseUserId(c.Param("uid"))
}
