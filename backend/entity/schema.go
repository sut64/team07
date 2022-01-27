package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID_student string `gorm:"uniqueIndex"`
	Prefix     string
	Name       string
	Major      string
	Year       uint
	Email      string
	Password   string

	Withdrawals  []Withdrawal  `gorm:"foreignKey:StudentID"`
	RequestExams []RequestExam `gorm:"foreignKey:StudentID"`
	//RecordPetition []RecordPetition `gorm:"foreignKey:StudentID"`
	//IncreaseGrades []IncreaseGrades `gorm:"foreignKey:StudentID"`
}

type Teacher struct {
	gorm.Model
	ID_teacher string `gorm:"uniqueIndex"`
	Prefix     string
	Name       string
	Major      string
	Email      string
	Password   string

	AddCourse []AddCourse `gorm:"foreignKey:TeacherID"`
}

type Registrar struct {
	gorm.Model
	ID_registrar string `gorm:"uniqueIndex"`
	Prefix       string
	Name         string
	Email        string
	Password     string
}

type Semester struct {
	gorm.Model
	Semester     string
	ExamSchedule []ExamSchedule `gorm:"foreignKey:SemesterID"`
}

type ExamType struct {
	gorm.Model
	Type         string
	ExamSchedule []ExamSchedule `gorm:"foreignKey:ExamTypeID"`
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

/*type RequestStatus struct {
	gorm.Model
	Status       string
	RequestExams []RequestExam `gorm:"foreignKey:RequestStatusID"`
}

type Grades struct {
	gorm.Model
	Grade string

	IncreaseGrades []IncreaseGrades `gorm:"foreignKey:GradesID"`
}

type Petition struct {
	gorm.Model
	Claim          string

	RecordPetition []RecordPetition `gorm:"foreignKey:PetitionID"`
}*/

type AddCourse struct {
	gorm.Model
	Credit   int16
	DayTime  string
	SaveTime time.Time

	CourseID *uint
	Course   Course `gorm:"references:id"`

	ProgramID *uint
	Program   Program `gorm:"references:id"`

	TeacherID *uint
	Teacher   Teacher `gorm:"references:id"`
}

type ExamSchedule struct {
	gorm.Model
	AcademicYear int16
	RoomExam     string
	DateExam     time.Time
	StartTime    time.Time
	EndTime      time.Time

	CourseID *uint
	Course   Course `gorm:"references:id"`

	ExamTypeID *uint
	ExamType   ExamType `gorm:"references:id"`

	SemesterID *uint
	Semester   Semester `gorm:"references:id"`
}

type Withdrawal struct {
	gorm.Model

	StudentID *uint
	Student   Student `gorm:"references:id" valid:"-"`

	CourseID *uint
	Course   Course `gorm:"references:id" valid:"-"`

	TeacherID *uint
	Teacher   Teacher `gorm:"references:id" valid:"-"`

	SemesterID *uint
	Semester   Semester `gorm:"references:id" valid:"-"`

	YearTime       int
	RemainCredit   int
	Reason         string
	WithdrawalTime time.Time
}

/*
	type RequestExam struct {
		gorm.Model

		StudentID *uint
		Student   Student `gorm:"references:id"`

		SemesterID *uint
		Semester   Semester `gorm:"references:id"`

		AcademicYear int

		CourseID *uint
		Course   Course `gorm:"references:id"`

		TeacherID *uint
		Teacher   Teacher `gorm:"references:id"`

		Tel string

		RequestStatusID *uint
		RequestStatus   RequestStatus `gorm:"references:id"`

		RequestTime time.Time
	}

	type RecordPetition struct {
		gorm.Model

		Because          string
		RegisteredCredit int
		TimeRecord       time.Time

		StudentID *uint
		Student   Student `gorm:"references:id"`

		PetitionID *uint
		Petition   Petition `gorm:"references:id"`

		CourseID *uint
		Course   Course `gorm:"references:id"`
	}


	type IncreaseGrades struct {
		gorm.Model
		Date   time.Time
		Credit uint

		StudentID *uint
		Student   Student `gorm:"references:id"`

		GradesID *uint
		Grades   Grades `gorm:"references:id"`

		CourseID *uint
		Course   Course `gorm:"references:id"`

	}*/
