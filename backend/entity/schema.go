package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
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

	RegisCourse    []RegisCourse    `gorm:"foreignKey:StudentID"`
	Withdrawals    []Withdrawal     `gorm:"foreignKey:StudentID"`
	RequestExams   []RequestExam    `gorm:"foreignKey:StudentID"`
	RecordPetition []RecordPetition `gorm:"foreignKey:StudentID"`
	IncreaseGrades []IncreaseGrades `gorm:"foreignKey:StudentID"`
}

type Teacher struct {
	gorm.Model
	ID_teacher string `gorm:"uniqueIndex"`
	Prefix     string
	Name       string
	Major      string
	Email      string
	Password   string

	AddCourse    []AddCourse   `gorm:"foreignKey:TeacherID"`
	RequestExams []RequestExam `gorm:"foreignKey:StudentID"`
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
	RequestExams []RequestExam  `gorm:"foreignKey:SemesterID"`
}

type ExamType struct {
	gorm.Model
	Type         string
	ExamSchedule []ExamSchedule `gorm:"foreignKey:ExamTypeID"`
}

type Course struct {
	gorm.Model
	Coursename   string
	Coursenumber int32 `gorm:"uniqueIndex"`

	RegisCourse    []RegisCourse    `gorm:"foreignKey:CourseID"`
	ExamSchedule   []ExamSchedule   `gorm:"foreignKey:CourseID"`
	AddCourse      []AddCourse      `gorm:"foreignKey:CourseID"`
	IncreaseGrades []IncreaseGrades `gorm:"foreignKey:CourseID"`
	RecordPetition []RecordPetition `gorm:"foreignKey:CourseID"`
	RequestExams   []RequestExam    `gorm:"foreignKey:CourseID"`
}

type Program struct {
	gorm.Model
	Programname string

	AddCourse []AddCourse `gorm:"foreignKey:ProgramID"`
}

type RequestStatus struct {
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
	Claim string

	RecordPetition []RecordPetition `gorm:"foreignKey:PetitionID"`
}

type AddCourse struct {
	gorm.Model
	Credit   int16     `valid:"range(1|4)~ข้อมูลหน่วยกิตผิดพลาด, required~ข้อมูลหน่วยกิตผิดพลาด"`
	DayTime  string    `valid:"required~ข้อมูลวันที่และเวลาที่สอนผิดพลาด"`
	SaveTime time.Time `valid:"DelayNow3Min~ข้อมูลวันที่และเวลาที่บันทึกผิดพลาด"`

	CourseID *uint
	Course   Course `gorm:"references:id" valid:"-"`

	ProgramID *uint
	Program   Program `gorm:"references:id" valid:"-"`

	TeacherID *uint
	Teacher   Teacher `gorm:"references:id" valid:"-"`
}

type ExamSchedule struct {
	gorm.Model
	AcademicYear int16     `valid:"range(2000|3000)~ปีการศึกษาต้องเป็นตัวเลข 4 หลัก, required~ปีการศึกษาต้องเป็นตัวเลข 4 หลัก"`
	RoomExam     string    `valid:"matches(^[B]\\d{4}$)~ห้องสอบต้องขึ้นต้นด้วย B และตามด้วยตัวเลข 4 ตัว, required~ห้องสอบต้องขึ้นต้นด้วย B และตามด้วยตัวเลข 4 ตัว"`
	ExamDate     time.Time `valid:"future~วันที่สอบต้องเป็นวันในอนาคต"`
	StartTime    time.Time
	EndTime      time.Time

	CourseID *uint
	Course   Course `gorm:"references:id" valid:"-"`

	ExamTypeID *uint
	ExamType   ExamType `gorm:"references:id" valid:"-"`

	SemesterID *uint
	Semester   Semester `gorm:"references:id" valid:"-"`
}

type RegisCourse struct {
	gorm.Model

	StudentID *uint
	Student   Student `gorm:"references:id"`

	CourseID *uint
	Course   Course `gorm:"references:coursenumber"`

	Withdrawal []Withdrawal `gorm:"foreignKey:RegisCourseID"`
}

type Withdrawal struct {
	gorm.Model

	StudentID *uint
	Student   Student `gorm:"references:id" valid:"-"`

	RegisCourseID *uint
	RegisCourse   RegisCourse `gorm:"references:id" valid:"-"`

	TeacherID *uint
	Teacher   Teacher `gorm:"references:id" valid:"-"`

	SemesterID *uint
	Semester   Semester `gorm:"references:id" valid:"-"`

	YearTime       int       `valid:"range(2000|3000)~ข้อมูลปีการศึกษาผิดพลาด, required~ข้อมูลปีการศึกษาผิดพลาด"`
	RemainCredit   int       `valid:"positive~ข้อมูลหน่วยกิตไม่ถูกต้อง, required~ข้อมูลหน่วยกิตไม่ถูกต้อง"`
	Reason         string    `valid:"required~ข้อมูลเหตุผลไม่ถูกต้อง"`
	WithdrawalTime time.Time `valid:"present~ข้อมูลวันเวลาไม่ถูกต้อง"`
}

type RequestExam struct {
	gorm.Model

	StudentID *uint
	Student   Student `gorm:"references:id" valid:"-"`

	SemesterID *uint
	Semester   Semester `gorm:"references:id" valid:"-"`

	AcademicYear int `valid:"required~ข้อมูลปีการศึกษาไม่ถูกต้อง,range(2500|2800)~ข้อมูลปีการศึกษาไม่ถูกต้อง"`

	CourseID *uint
	Course   Course `gorm:"references:id" valid:"-"`

	TeacherID *uint
	Teacher   Teacher `gorm:"references:id" valid:"-"`

	Tel string `valid:"required~ข้อมูลเบอร์ติดต่อไม่ถูกต้อง,matches(^0([6|8|9])([0-9]{8}$))~ข้อมูลเบอร์ติดต่อไม่ถูกต้อง"`

	RequestStatusID *uint
	RequestStatus   RequestStatus `gorm:"references:id"`

	RequestTime time.Time `valid:"present~ข้อมูลวันเวลาไม่ถูกต้อง"`
}

type RecordPetition struct {
	gorm.Model

	Because          string    `valid:"required~ข้อมูลเหตุผลไม่ถูกต้อง,stringlength(0|200)~ข้อมูลเหตุผลไม่ถูกต้อง"`
	RegisteredCredit int       `valid:"positive~ข้อมูลหน่วยกิตไม่ถูกต้อง, required~ข้อมูลหน่วยกิตไม่ถูกต้อง"`
	TimeRecord       time.Time `valid:"present~ข้อมูลวันเวลาไม่ถูกต้อง"`

	StudentID *uint
	Student   Student `gorm:"references:id" valid:"-"`

	PetitionID *uint
	Petition   Petition `gorm:"references:id" valid:"-"`

	CourseID *uint
	Course   Course `gorm:"references:id" valid:"-"`
}

type IncreaseGrades struct {
	gorm.Model
	Date        time.Time `valid:"present~ข้อมูลวันเวลาไม่ถูกต้อง"`
	GradePoint  int       `valid:"range(0|100)~ข้อมูลคะแนนไม่ถูกต้อง, required~ข้อมูลคะแนนไม่ถูกต้อง"`
	Description string    `valid:"length(0|50)~ข้อมูลหมายเหตุไม่ถูกต้อง"`

	StudentID *uint
	Student   Student `gorm:"references:id" valid:"-" `

	GradesID *uint
	Grades   Grades `gorm:"references:id" valid:"-"`

	CourseID *uint
	Course   Course `gorm:"references:id" valid:"-"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.After(t)
	})
	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.Before(time.Time(t))
	})
	govalidator.CustomTypeTagMap.Set("present", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return t.After(now.Add(3-time.Minute)) && t.Before(now.Add(3+time.Minute))
	})
	govalidator.CustomTypeTagMap.Set("DelayNow3Min", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().Add(3 - time.Minute))
	})
	govalidator.CustomTypeTagMap.Set("positive", func(i interface{}, context interface{}) bool {
		num := i
		return num.(int) > 0
	})

}
