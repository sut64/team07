package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestRequestExamPass(t *testing.T) {
	g := NewGomegaWithT(t)

	requestexam := RequestExam{
		AcademicYear: 2564,
		Tel:          "0941548155",
		RequestTime:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(requestexam)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

// ตรวจสอบปีการศึกษาต้องเป็นตัวเลขที่อยู่ในช่วง 2500-2600
func TestAcademicYearMustBeInRange(t *testing.T) {
	g := NewGomegaWithT(t)

	requestexam := RequestExam{
		AcademicYear: 2400, //ผิด
		Tel:          "0941548155",
		RequestTime:  time.Now(),
	}

	ok, err := govalidator.ValidateStruct(requestexam)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลปีการศึกษาไม่ถูกต้อง"))
}

// ตรวจสอบเบอร์ติดต่อต้องขึ้นต้นด้วย 0 ตามด้วย 6,8,9 และตามด้วย 0-9 จำนวน 8 ตัว
func TestTelMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	requestexam := RequestExam{
		AcademicYear: 2564,
		Tel:          "029058901", //ผิด
		RequestTime:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(requestexam)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลเบอร์ติดต่อไม่ถูกต้อง"))
}

// ตรวจสอบวันเวลาต้องไม่เป็นวันเวลาปัจจุบัน
func TestRequestTimeMustBePresent(t *testing.T) {
	g := NewGomegaWithT(t)

	requestexam := RequestExam{
		AcademicYear: 2564,
		Tel:          "0941548155",
		RequestTime:  time.Now().Add(24 * time.Hour), //ผิด
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(requestexam)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลวันเวลาไม่ถูกต้อง"))
}
