package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func createStudent(c *gin.Context) {
	var std transformedStudent
	var model student
	c.Bind(&std)
	validasi := validatorCreated(std)
	model = transferVoToModel(std)
	if validasi != "" {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": validasi})
	} else {
		db.Create(&model)
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": model})
	}
}

func fetchAllStudent(c *gin.Context) {
	var model [] student
	var vo [] transformedStudent

	db.Find(&model)

	if len(model) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
	}

	for _, item := range model {
		vo = append(vo, transferModelToVo(item))
	}
	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": vo})
}

func fetchSingleStudent(c *gin.Context) {
	var model student
	var vo transformedStudent

	modelID := c.Param("id")
	db.Find(&model, modelID)

	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
	}
	vo = transferModelToVo(model)
	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": vo})
}

func updateStudent(c *gin.Context) {
	var model student
	var vo transformedStudent
	modelID := c.Param("id")
	db.First(&model, modelID)

	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
	}
	c.Bind(&vo)

	validasi := validatorCreated(vo)
	if validasi != "" {
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": validasi})
	} else {
		db.Model(&model).Update(transferVoToModel(vo))
		c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": model})
	}
}

func deleteStudent(c *gin.Context) {
	var model student
	modelID := c.Param("id")

	db.First(&model, modelID)
	if model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Datat Tidak di Temukan"})
	}
	db.Delete(model)
	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": "Data Telah berhasil di hapus"})
}
