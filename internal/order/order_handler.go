package order

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	OrderService Service
}

func NewHandler(s Service) *handler {
	return &handler{
		OrderService: s,
	}
}

func (h *handler) CreateOrder(ctx *gin.Context) {
	var req CreateOrderRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	res, err := h.OrderService.CreateOrder(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Order has successfully created.",
		"data":    res,
	})
}

func (h *handler) GetOrder(ctx *gin.Context) {
	res, err := h.OrderService.GetOrder()
	if err != nil {
		if errors.Is(err, ErrOrderEmpty) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully retrieve orders",
		"data":    res,
	})
}

func (h *handler) DeleteOrder(ctx *gin.Context) {
	var req DeleteOrderRequest

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": ErrInvalidID.Error(),
		})
		return
	}
	req.Id = uint(id)

	err = h.OrderService.DeleteOrder(&req)
	if err != nil {
		if errors.Is(err, ErrNullRecordAffected) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "order not found",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully delete order",
	})
}

func (h *handler) UpdateOrder(ctx *gin.Context) {
	var req UpdateOrderRequest

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": ErrInvalidID.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	res, err := h.OrderService.UpdateOrder(uint(id), &req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "order not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Successfully delete order",
		"data":    res,
	})
}
