package entity

import (
	"testing"
	"time"
	//"fmt"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAddCoursePass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	addcourse := AddCourse{
		Credit:   4,
		DayTime:  "TUE 18.00-21.00",
		SaveTime: time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(addcourse)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestDayTimeNotBlank(t *testing.T) {
 	g := NewGomegaWithT(t)

 	addcourse := AddCourse{
		Credit:   4,
		DayTime:  "",//ผิด
		SaveTime: time.Now(),
	}

 	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(addcourse)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลวันที่และเวลาที่สอนผิดพลาด"))
 }

// ตรวจสอบวันเวลาที่บันทึกต้องไม่เป็นเวลาในอดีต
func TestSaveTimeMustNotBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	addcourse := AddCourse{
		Credit:   4,
		DayTime:  "TUE 18.00-21.00",
		SaveTime: time.Now().Add(5 - time.Hour), //ผิด
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(addcourse)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลวันที่และเวลาที่บันทึกผิดพลาด"))
}
// ตรวจสอบหน่วยกิตต้องเป็นตัวเลขที่อยู่ในช่วง 1-4
func TestCreditMustBeInRange(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []int16{
		10,
		-5,
	}
	for _, fixture := range fixtures{
		addcourse := AddCourse{
		Credit:   fixture, //ผิด
		DayTime:  "TUE 18.00-21.00",
		SaveTime: time.Now(),
		
	}
	
		ok, err := govalidator.ValidateStruct(addcourse)

		// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error() ต้องมี message แสดงออกมา
		g.Expect(err.Error()).To(Equal("ข้อมูลหน่วยกิตผิดพลาด"))
	}
}