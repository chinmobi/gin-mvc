// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	ctrl "github.com/chinmobi/gin-mvc/controller"

	"github.com/gin-gonic/gin"
)

func setupUserRoutes(r *gin.RouterGroup, userCtrl *ctrl.UserController) {
	updateUserFunc := updateUser(userCtrl)

	r.POST("/users",        createUser(userCtrl))
	r.GET("/users",         getAllUsers(userCtrl))
	r.GET("/users/:uid",    getUserByID(userCtrl))
	r.PATCH("/users/:uid",  updateUserFunc)
	r.PUT("/users/:uid",    updateUserFunc)
	r.DELETE("/users/:uid", deleteUser(userCtrl))
}

func createUser(userCtrl *ctrl.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		userCtrl.CreateUser(c)
	}
}

func getAllUsers(userCtrl *ctrl.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		userCtrl.GetAllUsers(c)
	}
}

func getUserByID(userCtrl *ctrl.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		userCtrl.GetUserByID(c)
	}
}

func updateUser(userCtrl *ctrl.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		userCtrl.UpdateUser(c)
	}
}

func deleteUser(userCtrl *ctrl.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		userCtrl.DeleteUser(c)
	}
}
