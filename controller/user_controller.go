package controller

import (
	"net/http"
	"strconv"
	"todo-api/models"
	"todo-api/models/dto"
	"todo-api/service"
	"todo-api/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
	rg          *gin.RouterGroup
}

func (t *UserController) Route() {
	router := t.rg.Group("/users")
	{
		router.GET("/", t.GetAllUsers)
		router.POST("/", t.CreateUser)
		router.PUT("/:id", t.UpdateUser)
		router.DELETE("/:id", t.DeleteUser)
		router.GET("/:id", t.GetUserByID)
	}
}

func NewUserController(userService service.UserService, rg *gin.RouterGroup) *UserController {
	return &UserController{
		userService: userService,
		rg:          rg,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	user := &models.User{
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	if err := c.userService.CreateUser(user); err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSingleResponse(ctx, "User created successfully", user, http.StatusCreated)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	user, err := c.userService.GetUserByID(userID)
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	if user == nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSingleResponse(ctx, "User retrieved successfully", user, http.StatusOK)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	size, err2 := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err2 != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	data, paging, err := c.userService.GetAllUsers(page, size)
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
	}

	var listData []interface{}

	for _, pl := range data {
		listData = append(listData, pl)
	}

	utils.SendPagingResponse(ctx, "Success Get Data", listData, paging, http.StatusOK)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	var req dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.userService.GetUserByID(userID)
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	if user == nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	user.Fullname = req.Fullname
	user.Email = req.Email
	user.Role = req.Role

	if err := c.userService.UpdateUser(user); err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSingleResponse(ctx, "User updated successfully", user, http.StatusOK)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	err := c.userService.DeleteUser(userID)
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSingleResponse(ctx, "User deleted successfully", nil, http.StatusOK)
}
