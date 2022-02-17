package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ข้อมูลถูกต้องหมดทุก field
func TestIncreaseGradesPass(t *testing.T) {
	g := NewGomegaWithT(t)

	increasegrades := IncreaseGrades{
		Date:        time.Now(),
		GradePoint:  100,
		Description: "so good",
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(increasegrades)

	// ok ต้องเป็น true แปลว่า ไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่า ไม่มี error
	g.Expect(err).To(BeNil())
}

// ตรวจสอบเวลาต้องไม่เป็นเวลาในอดีต
func TestDateMustBePresent(t *testing.T) {
	g := NewGomegaWithT(t)

	ig := IncreaseGrades{
		Date:        time.Now().Add(24 * time.Hour), //ผิด
		GradePoint:  100,
		Description: "so good",
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ig)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	//err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	//err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลวันเวลาไม่ถูกต้อง"))
}

//	ตรวจสอบคะแนนว่าอยู่ในช่วง 0 - 100
func TestGradePointMustBeInRange(t *testing.T) {
	g := NewGomegaWithT(t)

	ig := IncreaseGrades{
		Date:        time.Now(),
		GradePoint:  111, //ผิด
		Description: "so good",
	}

	ok, err := govalidator.ValidateStruct(ig)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	//err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	//err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลคะแนนไม่ถูกต้อง"))
}

//	ตรวจสอบหมายเหตุว่าอยู่ในช่วง
func TestDescriptionMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	ig := IncreaseGrades{
		Date:        time.Now(),
		GradePoint:  100,
		Description: "gooddddddddddddddddddddddddddddddddddddddddddddddddddddddddd", //ผิด
	}

	ok, err := govalidator.ValidateStruct(ig)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	//err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	//err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลหมายเหตุไม่ถูกต้อง"))
}
