package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetVersion() string {
	return "v1"
}

func (h *Handler) AddRoutes(r *gin.RouterGroup) {
	h.setUserRouterGroup(r)
	h.setTrackerRouterGroup(r)
}

func (h *Handler) setUserRouterGroup(r *gin.RouterGroup) {
	users := r.Group("/users")
	users.POST("/create", h.createUser)
	users.GET("/get", h.getUser)
	users.PUT("/update", h.updateUser)
	users.DELETE("/delete", h.deleteUser)
	users.GET("/get-by-filter", h.getUsers)
	users.GET("/get-by-filter-paged", h.getUsersPaged)
}

func (h *Handler) setTrackerRouterGroup(r *gin.RouterGroup) {
	tracker := r.Group("/tracker")
	tracker.GET("/get-user-costs-by-period", h.getUserCostsByPeriod)
	tracker.POST("/start", h.start)
	tracker.PUT("/stop", h.stop)
}
