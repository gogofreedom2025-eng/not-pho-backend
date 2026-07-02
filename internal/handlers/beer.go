package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gogo1/not-pho-backend/internal/models"
	"gorm.io/gorm"
)

type BeerHandler struct {
	DB *gorm.DB
}

type createBeerRequest struct {
	Name        string  `json:"name" binding:"required"`
	Brewery     string  `json:"brewery" binding:"required"`
	Style       string  `json:"style" binding:"required"`
	ABV         float64 `json:"abv" binding:"required,gt=0,lte=20"`
	Description string  `json:"description"`
}

type updateBeerRequest struct {
	Name        string  `json:"name" binding:"required"`
	Brewery     string  `json:"brewery" binding:"required"`
	Style       string  `json:"style" binding:"required"`
	ABV         float64 `json:"abv" binding:"required,gt=0,lte=20"`
	Description string  `json:"description"`
}

func (h *BeerHandler) List(c *gin.Context) {
	var beers []models.Beer
	if err := h.DB.Order("id asc").Find(&beers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list beers"})
		return
	}

	c.JSON(http.StatusOK, beers)
}

func (h *BeerHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var beer models.Beer
	if err := h.DB.First(&beer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "beer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get beer"})
		return
	}

	c.JSON(http.StatusOK, beer)
}

func (h *BeerHandler) Create(c *gin.Context) {
	var req createBeerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	beer := models.Beer{
		Name:        req.Name,
		Brewery:     req.Brewery,
		Style:       req.Style,
		ABV:         req.ABV,
		Description: req.Description,
	}

	if err := h.DB.Create(&beer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create beer"})
		return
	}

	c.JSON(http.StatusCreated, beer)
}

func (h *BeerHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req updateBeerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var beer models.Beer
	if err := h.DB.First(&beer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "beer not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get beer"})
		return
	}

	beer.Name = req.Name
	beer.Brewery = req.Brewery
	beer.Style = req.Style
	beer.ABV = req.ABV
	beer.Description = req.Description

	if err := h.DB.Save(&beer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update beer"})
		return
	}

	c.JSON(http.StatusOK, beer)
}

func (h *BeerHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	result := h.DB.Delete(&models.Beer{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete beer"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "beer not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
