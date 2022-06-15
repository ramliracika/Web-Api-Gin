package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramliracika/Web-Api-Go/laptop"
	"github.com/ramliracika/Web-Api-Go/service"
)

type Laptophandler struct {
	laptopService service.Service
}

func NewHandler(laptopService service.Service) *Laptophandler {
	return &Laptophandler{laptopService}
}

func (h *Laptophandler) MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  "Ramli",
		"dream": "As Software Engineer ASAP",
	})
}

func (h *Laptophandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Content": "Hello Server",
	})
}

func (h *Laptophandler) ParamsHandler(c *gin.Context) {
	id := c.Param("id") //get parameter like /laptop/id
	brand := c.Param("brand")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"brand": brand,
	})
}

func (h *Laptophandler) QueryHandler(c *gin.Context) {
	stock := c.Query("stock") // get query like ?stock=
	location := c.Query("location")
	c.JSON(http.StatusOK, gin.H{
		"stock":    stock,
		"location": location,
	})
}

func (h *Laptophandler) PostHandler(c *gin.Context) {
	var laptop laptop.Gaming
	err := c.ShouldBindJSON(&laptop)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return

	}
	NewLaptop, err := h.laptopService.Create(laptop)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": NewLaptop,
	})
}

func (h *Laptophandler) GetAllHandler(c *gin.Context) {
	laptop, err := h.laptopService.FindAll()
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": laptop,
	})
}

func (h *Laptophandler) GetByIdHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	laptop, err := h.laptopService.FindById(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if laptop.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": "Data Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": laptop,
	})

}

func (h *Laptophandler) GetByIdQueryHandler(c *gin.Context) {
	id := c.Query("id")
	ID, _ := strconv.Atoi(id)
	laptop, err := h.laptopService.FindById(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if laptop.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": "Data Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": laptop,
	})

}

func (h *Laptophandler) UpdateHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	var laptop laptop.Gaming
	laptop, err := h.laptopService.FindById(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	if laptop.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": "Data Not Found",
		})
		return
	}

	c.ShouldBindJSON(&laptop)

	UpdateLaptop, err := h.laptopService.Update(laptop, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Messages": "Data succeed to update",
		"data":     UpdateLaptop,
	})

}

func (h *Laptophandler) DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)

	var laptop laptop.Gaming
	laptop, err := h.laptopService.FindById(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	if laptop.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data": "Data Not Found",
		})
		return
	}

	_, err = h.laptopService.Delete(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messeges": "Succesfully deleted from database",
	})
}
