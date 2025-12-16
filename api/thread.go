package api

import (
	"net/http"
	"strconv"

	"github.com/Iknite-Space/sqlc-example-api/db/repo"
	"github.com/gin-gonic/gin"
)

type ThreadHandler struct {
	repo *repo.Queries
}

func NewThreadHandler(r *repo.Queries) *ThreadHandler {
	return &ThreadHandler{repo: r}
}

func (h *ThreadHandler) WireHttpHandler(router *gin.Engine) {
	router.POST("/thread", h.handleCreateThread)
	router.GET("/thread/:id", h.handleGetThread)
	router.GET("/threads", h.handleListThreads)
	router.GET("/thread/:id/messages", h.handleGetMessagesByThread)
}

// Create Thread
func (h *ThreadHandler) handleCreateThread(c *gin.Context) {
	var req struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thread, err := h.repo.CreateThread(c.Request.Context(), req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)
}

// Get Thread by ID
func (h *ThreadHandler) handleGetThread(c *gin.Context) {
	id := c.Param("id")
	thread, err := h.repo.GetThread(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "thread not found"})
		return
	}
	c.JSON(http.StatusOK, thread)
}

// List Threads
func (h *ThreadHandler) handleListThreads(c *gin.Context) {
	threads, err := h.repo.ListThreads(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, threads)
}

// // Paginated Messages by Thread
// func (h *ThreadHandler) handleGetMessagesByThread(c *gin.Context) {
// 	thread := c.Param("id")

// 	// Simple pagination parameters: start & count
// 	startStr := c.DefaultQuery("start", "0")
// 	countStr := c.DefaultQuery("count", "10")

// 	start, err := strconv.Atoi(startStr)
// 	if err != nil || start < 0 {
// 		start = 0
// 	}

// 	count, err := strconv.Atoi(countStr)
// 	if err != nil || count < 1 {
// 		count = 10
// 	}

// 	// Use the correct fields from generated sqlc code
// 	messages, err := h.repo.GetMessagesByThreadPaginated(
// 		c.Request.Context(),
// 		repo.GetMessagesByThreadPaginatedParams{
// 			Thread:  thread,       // matches "thread" column
// 			Column2: int32(count), // LIMIT
// 			Column3: int32(start), // OFFSET
// 		},
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

//		c.JSON(http.StatusOK, messages)
//	}
//
// Paginated Messages
func (h *ThreadHandler) handleGetMessagesByThread(c *gin.Context) {
	thread := c.Param("id")

	// Pagination parameters
	startStr := c.DefaultQuery("start", "0")
	countStr := c.DefaultQuery("count", "10")

	start, err := strconv.Atoi(startStr)
	if err != nil || start < 0 {
		start = 0
	}

	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		count = 10
	}

	messages, err := h.repo.GetMessagesByThreadPaginated(
		c.Request.Context(),
		repo.GetMessagesByThreadPaginatedParams{
			Thread:  thread,
			Column2: int32(count),
			Column3: int32(start),
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
