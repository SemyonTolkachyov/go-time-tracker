package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-time-tracker/internal/apperror"
	httpmodel "go-time-tracker/internal/model/http"
	"go-time-tracker/internal/model/input"
	"go-time-tracker/internal/utils"
	"net/http"
	"strconv"
)

// @Summary Get user costs
// @Tags tracker
// @Description get user costs by period
// @ID get-user-costs
// @Accept  json
// @Produce  json
// @Param userId query int true "userId"
// @Param start query string true "start"
// @Param end query string true "end"
// @Success 200 {object} http.StatusResponse
// @Failure 400,404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Failure default {object} http.ErrorResponse
// @Router /api/v1/tracker/get-user-costs-by-period [get]
func (h *Handler) getUserCostsByPeriod(c *gin.Context) {
	params := c.Request.URL.Query()
	userId, err := strconv.Atoi(params.Get("userId"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}
	start, err := utils.ParseAsTime(params.Get("start"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid start date value")
	}

	end, err := utils.ParseAsTime(params.Get("end"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid end date value")
	}

	i := input.TimeCostsByPeriod{
		UserId: userId,
		Start:  start,
		End:    end,
	}

	costs, err := h.services.TrackerService.GetUserCostsByPeriod(c, i)
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(*costs) == 0 {
		httpmodel.NewErrorResponse(c, http.StatusNotFound, "user costs not found")
		return
	}

	c.JSON(http.StatusOK, costs)
}

// @Summary Start
// @Tags tracker
// @Description start timer
// @ID start
// @Accept  json
// @Produce  json
// @Param userId query int true "userId"
// @Param taskId query int true "taskId"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Failure default {object} http.ErrorResponse
// @Router /api/v1/tracker/start [post]
func (h *Handler) start(c *gin.Context) {
	params := c.Request.URL.Query()
	userId, err := strconv.Atoi(params.Get("userId"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}
	taskId, err := strconv.Atoi(params.Get("taskId"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid task id")
		return
	}
	err = h.services.TrackerService.Start(c, userId, taskId)
	if err != nil {
		if errors.Is(err, apperror.ExistsErr) || errors.Is(err, apperror.BadRequestErr) {
			httpmodel.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		httpmodel.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, httpmodel.StatusResponse{Status: "ok"})
	return
}

// @Summary Stop
// @Tags tracker
// @Description stop timer
// @ID stop
// @Accept  json
// @Produce  json
// @Param userId query int true "userId"
// @Param taskId query int true "taskId"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} http.ErrorResponse
// @Failure 500 {object} http.ErrorResponse
// @Failure default {object} http.ErrorResponse
// @Router /api/v1/tracker/stop [put]
func (h *Handler) stop(c *gin.Context) {
	params := c.Request.URL.Query()
	userId, err := strconv.Atoi(params.Get("userId"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}
	taskId, err := strconv.Atoi(params.Get("taskId"))
	if err != nil {
		httpmodel.NewErrorResponse(c, http.StatusBadRequest, "invalid task id")
		return
	}
	err = h.services.TrackerService.Stop(c, userId, taskId)
	if err != nil {
		if errors.Is(err, apperror.NotFoundErr) {
			httpmodel.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		httpmodel.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, httpmodel.StatusResponse{Status: "ok"})
	return
}
