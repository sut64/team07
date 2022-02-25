package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team07/entity"
)

// POST /regiscourses
func CreateRegisCourse(c *gin.Context) {
	var RegisCourse entity.RegisCourse
	if err := c.ShouldBindJSON(&RegisCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&RegisCourse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": RegisCourse})
}

// GET /regiscourses/:id
func GetRegisCourse(c *gin.Context) {
	var RegisCourse entity.RegisCourse
	id := c.Param("id")
	if err := entity.DB().Preload("Student").Raw("SELECT * FROM regis_courses WHERE id = ?", id).Scan(&RegisCourse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RegisCourse})
}

// GET /regiscourses
func ListRegisCoursess(c *gin.Context) {
	var RegisCourses []entity.RegisCourse
	if err := entity.DB().Preload("Course").Raw("SELECT * FROM regis_courses ").Scan(&RegisCourses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RegisCourses})
}

func ListRegisCourses(c *gin.Context) {
	var RegisCourses []entity.RegisCourse
	id := c.Param("id")
	if err := entity.DB().Preload("Student").Preload("Course").Raw("SELECT * FROM regis_courses where student_id = ? ", id).Scan(&RegisCourses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RegisCourses})
}

// DELETE /regiscourses/:id
func DeleteRegisCourse(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM regis_courses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RegisCourse not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /regiscourses
func UpdateRegisCourse(c *gin.Context) {
	var RegisCourse entity.Course
	if err := c.ShouldBindJSON(&RegisCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", RegisCourse.ID).First(&RegisCourse); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RegisCourse not found"})
		return
	}

	if err := entity.DB().Save(&RegisCourse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": RegisCourse})
}
