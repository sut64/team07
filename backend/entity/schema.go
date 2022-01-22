package entity
import (
	  "time"
	  "gorm.io/gorm"
	)
	 
	type Semester struct {
		gorm.Model
		Semester  string
		ExamSchedule []ExamSchedule `gorm:"foreignKey:SemesterID"`
	}

	type ExamType struct {
		gorm.Model
		Type  string
		ExamSchedule []ExamSchedule `gorm:"foreignKey:ExamTypeID"`
	}

	type ExamSchedule struct {
		gorm.Model
		AcamedicYear int16
		RoomExam     string
		DateExam	 time.Time
		StartTime 	 time.Time
		EndTime		 time.Time
	
		/*CourseID *uint
		Course   Course `gorm:"references:id"`*/
	
		ExamTypeID *uint
		ExamType   ExamType `gorm:"references:id"`
	
		SemesterID *uint
		Semester   Semester `gorm:"references:id"`
	}