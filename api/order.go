// // // package api

// // // import (
// // // 	"net/http"

// // // 	"github.com/Iknite-Space/sqlc-example-api/db/repo"
// // // 	"github.com/gin-gonic/gin"
// // // )

// // // type CampayClient struct {
// // // 	Token string
// // // }

// // // type CampayRequest struct {
// // // 	Amount      int32
// // // 	Currency    string
// // // 	From        string
// // // 	Description string
// // // }

// // // func (c *CampayClient) CollectPayment(req CampayRequest) error {
// // // 	// Here you would call Campay API
// // // 	// For now, just log
// // // 	println("Campay payment requested:", req.Amount, req.Currency, req.From)
// // // 	return nil
// // // }

// // // type OrderHandler struct {
// // // 	querier repo.Querier
// // // 	campay  *CampayClient
// // // }

// // // func NewOrderHandler(q repo.Querier, campay *CampayClient) *OrderHandler {
// // // 	return &OrderHandler{
// // // 		querier: q,
// // // 		campay:  campay,
// // // 	}
// // // }

// // // func (h *OrderHandler) WireRoutes(r *gin.Engine) {
// // // 	r.POST("/orders", h.createOrder)
// // // }

// // // func (h *OrderHandler) createOrder(c *gin.Context) {
// // // 	var req struct {
// // // 		Name   string `json:"name" binding:"required"`
// // // 		Phone  string `json:"phone" binding:"required"`
// // // 		Amount int32  `json:"amount" binding:"required"`
// // // 	}

// // // 	if err := c.ShouldBindJSON(&req); err != nil {
// // // 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// // // 		return
// // // 	}

// // // 	order, err := h.querier.CreateOrder(c.Request.Context(), repo.CreateOrderParams{
// // // 		CustomerName:  req.Name,
// // // 		CustomerPhone: req.Phone,
// // // 		TotalAmount:   req.Amount,
// // // 		Currency:      "XAF",
// // // 		Status:        "PENDING",
// // // 	})
// // // 	if err != nil {
// // // 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// // // 		return
// // // 	}

// // // 	// Send payment request
// // // 	_ = h.campay.CollectPayment(CampayRequest{
// // // 		Amount:      req.Amount,
// // // 		Currency:    "XAF",
// // // 		From:        req.Phone,
// // // 		Description: "Order payment",
// // // 	})

// // // 	c.JSON(http.StatusCreated, order)
// // // }

// // package api

// // import (
// // 	"net/http"

// // 	"github.com/Iknite-Space/sqlc-example-api/db/repo"
// // 	"github.com/gin-gonic/gin"
// // )

// // type OrderHandler struct {
// // 	querier repo.Querier
// // 	campay  *CampayClient
// // }

// // func NewOrderHandler(q repo.Querier, campay *CampayClient) *OrderHandler {
// // 	return &OrderHandler{
// // 		querier: q,
// // 		campay:  campay,
// // 	}
// // }

// // func (h *OrderHandler) WireRoutes(r *gin.Engine) {
// // 	r.POST("/orders", h.createOrder)
// // }

// // func (h *OrderHandler) createOrder(c *gin.Context) {
// // 	var req struct {
// // 		Name   string `json:"name" binding:"required"`
// // 		Phone  string `json:"phone" binding:"required"`
// // 		Amount int32  `json:"amount" binding:"required"`
// // 	}

// // 	if err := c.ShouldBindJSON(&req); err != nil {
// // 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// // 		return
// // 	}

// // 	order, err := h.querier.CreateOrder(c.Request.Context(), repo.CreateOrderParams{
// // 		CustomerName:  req.Name,
// // 		CustomerPhone: req.Phone,
// // 		TotalAmount:   req.Amount,
// // 		Currency:      "XAF",
// // 		Status:        "PENDING",
// // 	})
// // 	if err != nil {
// // 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// // 		return
// // 	}

// // 	// Call Campay (convert int32 -> int)
// // 	_ = h.campay.CollectPayment(CampayRequest{
// // 		Amount:      int(req.Amount),
// // 		Currency:    "XAF",
// // 		From:        req.Phone,
// // 		Description: "Order payment",
// // 	})

// // 	c.JSON(http.StatusCreated, order)
// // }

// package api

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/Iknite-Space/sqlc-example-api/db/repo"
// 	"github.com/gin-gonic/gin"
// )

// type OrderHandler struct {
// 	querier repo.Querier
// 	campay  *CampayClient
// }

// func NewOrderHandler(q repo.Querier) *OrderHandler {
// 	// Initialize CampayClient here with your token
// 	campayClient := &CampayClient{
// 		Token: "62499bc353a726d8e4a6688dc46b44781e987b18", // YOUR CAMPAY TOKEN
// 	}

// 	return &OrderHandler{
// 		querier: q,
// 		campay:  campayClient,
// 	}
// }

// func (h *OrderHandler) WireRoutes(r *gin.Engine) {
// 	r.POST("/orders", h.createOrder)
// }

// func (h *OrderHandler) createOrder(c *gin.Context) {
// 	var req struct {
// 		Name   string `json:"name" binding:"required"`
// 		Phone  string `json:"phone" binding:"required"`
// 		Amount int32  `json:"amount" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Save order in database
// 	order, err := h.querier.CreateOrder(c.Request.Context(), repo.CreateOrderParams{
// 		CustomerName:  req.Name,
// 		CustomerPhone: req.Phone,
// 		TotalAmount:   req.Amount,
// 		Currency:      "XAF",
// 		Status:        "PENDING",
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Send payment request via Campay
// 	err = h.campay.CollectPayment(CampayRequest{
// 		Amount:      int(req.Amount),
// 		Currency:    "XAF",
// 		From:        "237680661612", // YOUR PHONE NUMBER
// 		Description: "Payment for Order #" + order.ID.String(),
// 	})
// 	if err != nil {
// 		fmt.Println("❌ Campay payment error:", err)
// 	}

// 	c.JSON(http.StatusCreated, order)
// }

package api

import (
	"fmt"
	"net/http"

	"github.com/Iknite-Space/sqlc-example-api/db/repo"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	querier repo.Querier
	campay  *CampayClient
}

// NewOrderHandler initializes the handler with a CampayClient
func NewOrderHandler(q repo.Querier) *OrderHandler {
	campayClient := &CampayClient{
		Token: "62499bc353a726d8e4a6688dc46b44781e987b18", // YOUR CAMPAY TOKEN
	}

	return &OrderHandler{
		querier: q,
		campay:  campayClient,
	}
}

func (h *OrderHandler) WireRoutes(r *gin.Engine) {
	r.POST("/orders", h.createOrder)
}

func (h *OrderHandler) createOrder(c *gin.Context) {
	var req struct {
		Name   string `json:"name" binding:"required"`
		Phone  string `json:"phone" binding:"required"`
		Amount int32  `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.querier.CreateOrder(c.Request.Context(), repo.CreateOrderParams{
		CustomerName:  req.Name,
		CustomerPhone: req.Phone,
		TotalAmount:   req.Amount,
		Currency:      "XAF",
		Status:        "PENDING",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trigger Campay payment
	err = h.campay.CollectPayment(CampayRequest{
		Amount:      int(req.Amount),
		Currency:    "XAF",
		From:        "237680661612", // YOUR PHONE NUMBER
		Description: "Payment for Order #" + order.ID.String(),
	})
	if err != nil {
		fmt.Println("❌ Campay payment error:", err)
	}

	c.JSON(http.StatusCreated, order)
}
