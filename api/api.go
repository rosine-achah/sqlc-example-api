package api

import (
	"net/http"
	"strconv"

	"github.com/Iknite-Space/sqlc-example-api/db/repo"
	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	repo *repo.Queries
}

func NewMessageHandler(r *repo.Queries) *MessageHandler {
	return &MessageHandler{repo: r}
}

func (h *MessageHandler) WireHttpHandler(r *gin.Engine) {
	// Thread endpoints
	r.POST("/thread", h.handleCreateThread)
	r.GET("/thread/:id", h.handleGetThread)

	// Message endpoints
	r.POST("/message", h.handleCreateMessage)
	r.GET("/message/:id", h.handleGetMessage)
	r.GET("/thread/:id/messages", h.handleGetMessagesByThread)
	r.PUT("/message/:id", h.handleUpdateMessage)
	r.DELETE("/message/:id", h.handleDeleteMessage)
}

// --- Thread Handlers ---
func (h *MessageHandler) handleCreateThread(c *gin.Context) {
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

func (h *MessageHandler) handleGetThread(c *gin.Context) {
	id := c.Param("id")
	thread, err := h.repo.GetThread(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "thread not found"})
		return
	}
	c.JSON(http.StatusOK, thread)
}

// --- Message Handlers ---
func (h *MessageHandler) handleCreateMessage(c *gin.Context) {
	var req struct {
		ThreadID string `json:"thread_id" binding:"required"`
		Sender   string `json:"sender" binding:"required"`
		Content  string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.repo.GetThread(c.Request.Context(), req.ThreadID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thread does not exist"})
		return
	}

	msg, err := h.repo.CreateMessage(c.Request.Context(), repo.CreateMessageParams{
		Thread:  req.ThreadID, // matches table column "thread"
		Sender:  req.Sender,
		Content: req.Content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (h *MessageHandler) handleGetMessage(c *gin.Context) {
	id := c.Param("id")
	msg, err := h.repo.GetMessageByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "message not found"})
		return
	}
	c.JSON(http.StatusOK, msg)
}

// Paginated messages by thread
func (h *MessageHandler) handleGetMessagesByThread(c *gin.Context) {
	thread := c.Param("id")

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

	msgs, err := h.repo.GetMessagesByThreadPaginated(
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

	c.JSON(http.StatusOK, msgs)
}

func (h *MessageHandler) handleUpdateMessage(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.repo.UpdateMessageContent(c.Request.Context(), repo.UpdateMessageContentParams{
		ID:      id,
		Content: req.Content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (h *MessageHandler) handleDeleteMessage(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeleteMessageByID(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
