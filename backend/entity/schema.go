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
	type Course struct {
		gorm.Model
		Coursename   string
		Coursenumber int32
	
		ExamSchedule []ExamSchedule `gorm:"foreignKey:CourseID"`
		AddCourse    []AddCourse    `gorm:"foreignKey:CourseID"`
	}

	type Program struct {
		gorm.Model
		Programname string
	
		AddCourse []AddCourse `gorm:"foreignKey:ProgramID"`
	}