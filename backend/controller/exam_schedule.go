package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team07/entity"
)

// POST /ExamSchedules
func CreateExamSchedule(c *gin.Context) {

	var ExamSchedule entity.ExamSchedule
	var Semester entity.Semester
	var ExamType entity.ExamType
	var Course entity.Course

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร ExamSchedule
	if err := c.ShouldBindJSON(&ExamSchedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// 9: ค้นหา Semester ด้วย id
	if tx := entity.DB().Where("id = ?", ExamSchedule.SemesterID).First(&Semester); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Semester not found"})
		return
	}

	// 10: ค้นหา ExamType ด้วย id
	if tx := entity.DB().Where("id = ?", ExamSchedule.ExamTypeID).First(&ExamType); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ExamType not found"})
		return
	}

	// 11: ค้นหา Course ด้วย id
	if tx := entity.DB().Where("id = ?", ExamSchedule.CourseID).First(&Course); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}

	// 12: สร้าง ExamSchedule
	es := entity.ExamSchedule{
		Semester: Semester,       // โยงความสัมพันธ์กับ Entity Semester
		Course:    Course,     // โยงความสัมพันธ์กับ Entity Course
		ExamType:  ExamType,        // โยงความสัมพันธ์กับ Entity ExamType
		AcademicYear: ExamSchedule.AcademicYear,
		ExamRoom: ExamSchedule.ExamRoom,
		ExamDate: ExamSchedule.ExamDate,
		StartTime: ExamSchedule.StartTime,
		EndTime: ExamSchedule.EndTime,
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(es); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&es).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": es})
}

// GET /ExamSchedule/:id
func GetExamSchedule(c *gin.Context) {
	var ExamSchedule entity.ExamSchedule
	id := c.Param("id")
	if err := entity.DB().Preload("Semester").Preload("ExamType").Preload("Course").Raw("SELECT * FROM exam_schedules WHERE id = ?", id).Find(&ExamSchedule).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ExamSchedule})
}

// GET /ExamSchedules
func ListExamSchedules(c *gin.Context) {
	var ExamSchedules []entity.ExamSchedule
	if err := entity.DB().Preload("Semester").Preload("ExamType").Preload("Course").Raw("SELECT * FROM exam_schedules").Find(&ExamSchedules).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ExamSchedules})
}

// DELETE /ExamSchedules/:id
func DeleteExamSchedule(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM exam_schedules WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ExamSchedule not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /ExamSchedules
func UpdateExamSchedule(c *gin.Context) {
	var ExamSchedule entity.ExamSchedule
	if err := c.ShouldBindJSON(&ExamSchedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", ExamSchedule.ID).First(&ExamSchedule); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ExamSchedule not found"})
		return
	}

	if err := entity.DB().Save(&ExamSchedule).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ExamSchedule})
}
