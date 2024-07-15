package v1

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"go-time-tracker/internal/apperror"
	"go-time-tracker/internal/model/filter"
	httpmodel "go-time-tracker/internal/model/http"
	"go-time-tracker/internal/model/input"
	"net/http"
	"strconv"
)

// @Summary Create user
// @Tags users
// @Description create user
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body input.NewUser true "passport info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Failure default {object} http.ErrorResponse
// @Router /api/v1/users/create [post]
func (h *Handler) createUser(c *gin.Context) {
	var i input.NewUser

	if err := c.BindJSON(&i); err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.UserService.Create(c, i)
	if err != nil {
		if errors.Is(err, apperror.ParseErr) {
			httpmodel.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		httpmodel.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get user
// @Tags users
// @Description get user by id
// @ID get-user-by-id
// @Produce  json
// @Param id query integer true "id"
// @Success 200 {object} http.StatusResponse
// @Failure 400,404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Failure default {object} http.ErrorResponse
// @Router /api/v1/users/get [get]
func (h *Handler) getUser(c *gin.Context) {
	params := c.Request.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	user, err := h.services.UserService.Get(c, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpmodel.NewErrorResponse(c, http.StatusNotFound, "user not found")
			return
		}
		httpmodel.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Update user
// @Tags users
// @Description update user
// @ID update-user
// @Accept  json
// @Produce  json
// @Param input body input.UpdateUser true "user info"
// @Success 200 {object} http.StatusResponse
// @Failure 400,404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Failure default {object} http.ErrorResponse
// @Router /api/v1/users/update [put]
func (h *Handler) updateUser(c *gin.Context) {
	var inp input.UpdateUser

	if err := c.BindJSON(&inp); err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	i, err := h.services.UserService.Update(c, inp)
	if err != nil {
		if errors.Is(err, apperror.ValidationErr) {
			httpmodel.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		httpmodel.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if i == 0 {
		httpmodel.NewErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}
	c.JSON(http.StatusOK, httpmodel.StatusResponse{Status: "ok"})
}

// @Summary Delete user
// @Tags users
// @Description delete user by id
// @ID delete-user-by-id
// @Produce  json
// @Param id query integer true "id"
// @Success 200 {object} http.StatusResponse
// @Failure 400,404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Failure default {object} http.ErrorResponse
// @Router /api/v1/users/delete [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	params := c.Request.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	i, err := h.services.UserService.Delete(c, id)
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if i == 0 {
		httpmodel.NewErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}
	c.JSON(http.StatusOK, httpmodel.StatusResponse{Status: "ok"})
}

// @Summary Get users
// @Tags users
// @Description get user by filter
// @ID get-user-by-filter
// @Produce  json
// @Param id query integer false "id"
// @Param createdAt query string false "createdAt"
// @Param updatedAt query string false "updatedAt"
// @Param passport query string false "passport"
// @Param surname query string false "surname"
// @Param name query string false "name"
// @Param patronymic query string false "patronymic"
// @Param address query string false "address"
// @Success 200 {string} string
// @Failure 400,404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Failure default {object} http.ErrorResponse
// @Router /api/v1/users/get-by-filter [get]
func (h *Handler) getUsers(c *gin.Context) {
	queryFilter, err := h.parseFilter(c)
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	users, err := h.services.UserService.GetByFilter(c, queryFilter)
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if users == nil {
		httpmodel.NewErrorResponse(c, http.StatusNotFound, "users not found")
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Get users paged
// @Tags users
// @Description get user by filter paged
// @ID get-user-by-filter-paged
// @Produce  json
// @Param size query integer false "size"
// @Param number query integer false "number"
// @Param id query integer false "id"
// @Param createdAt query string false "createdAt"
// @Param updatedAt query string false "updatedAt"
// @Param passport query string false "passport"
// @Param surname query string false "surname"
// @Param name query string false "name"
// @Param patronymic query string false "patronymic"
// @Param address query string false "address"
// @Success 200 {string} string
// @Failure 400,404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Failure default {object} http.ErrorResponse
// @Router /api/v1/users/get-by-filter-paged [get]
func (h *Handler) getUsersPaged(c *gin.Context) {
	params := c.Request.URL.Query()
	size, err := strconv.Atoi(params.Get("size"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid page size")
		return
	}
	number, err := strconv.Atoi(params.Get("number"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid page number")
		return
	}

	queryFilter, err := h.parseFilter(c)
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	users, err := h.services.UserService.GetPagedByFilter(c, size, number, queryFilter)
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if users == nil {
		httpmodel.NewErrorResponse(c, http.StatusNotFound, "users not found")
		return
	}

	c.JSON(http.StatusOK, users)
}

// parseFilter parse filter from query
func (h *Handler) parseFilter(c *gin.Context) (filter.User, error) {
	params := c.Request.URL.Query()
	userFilter := filter.User{}

	if idParam := params.Get("id"); idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return userFilter, err
		}
		userFilter.Id = new(int)
		*userFilter.Id = id
	}
	if createdAtParam := params.Get("createdAt"); createdAtParam != "" {
		createdAtFilter, err := filter.ParseTimeFilter(createdAtParam)
		if err != nil {
			return userFilter, err
		}
		userFilter.CreatedAt = new(filter.TimeFilter)
		userFilter.CreatedAt = createdAtFilter
	}
	if updatedAtParam := params.Get("updatedAt"); updatedAtParam != "" {
		updatedAtFilter, err := filter.ParseTimeFilter(updatedAtParam)
		if err != nil {
			return userFilter, err
		}
		userFilter.UpdatedAt = new(filter.TimeFilter)
		userFilter.UpdatedAt = updatedAtFilter
	}
	if passportParam := params.Get("passport"); passportParam != "" {
		userFilter.PassportNumber = new(string)
		*userFilter.PassportNumber = passportParam
	}
	if surnameParam := params.Get("surname"); surnameParam != "" {
		userFilter.Surname = new(string)
		*userFilter.Surname = surnameParam
	}
	if nameParam := params.Get("name"); nameParam != "" {
		userFilter.Name = new(string)
		*userFilter.Name = nameParam
	}
	if patronymicParam := params.Get("patronymic"); patronymicParam != "" {
		userFilter.Patronymic = new(string)
		*userFilter.Patronymic = patronymicParam
	}
	if addressParam := params.Get("address"); addressParam != "" {
		userFilter.Address = new(string)
		*userFilter.Address = addressParam
	}

	return userFilter, nil
}
