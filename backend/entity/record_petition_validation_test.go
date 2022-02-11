package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestRecordPetitionPass(t *testing.T) {
	g := NewGomegaWithT(t)

	recordpetition := RecordPetition{
		Because:          "ต้องการเรียนจบไวขึ้น",
		RegisteredCredit: 24,
		TimeRecord:       time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(recordpetition)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

// ตรวจสอบเหตุผลที่ต้องการลงทะเบียนเรียนเกิน/ต่ำกว่าหน่วยกิตที่กำหนด ตัวอักษรไม่เกิน 200 ตัวอักษร
func TestBecauseMustBeInRange(t *testing.T) {
	g := NewGomegaWithT(t)

	recordpetition := RecordPetition{

		Because:          "", //ผิด
		RegisteredCredit: 24,
		TimeRecord:       time.Now(),
	}

	ok, err := govalidator.ValidateStruct(recordpetition)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลเหตุผลไม่ถูกต้อง"))
}

// ตรวจสอบเบอร์หน่วยกิตต้องเป็นตัวเลขและไม่ติดลบ
func TestRegisteredCreditMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	recordpetition := RecordPetition{
		Because:          "ต้องการเรียนจบไวขึ้น",
		RegisteredCredit: -1, //ผิด
		TimeRecord:       time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(recordpetition)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลหน่วยกิตไม่ถูกต้อง"))
}

// ตรวจสอบวันเวลาต้องไม่เป็นวันเวลาปัจจุบัน
func TestRecordTimeMustBePresent(t *testing.T) {
	g := NewGomegaWithT(t)

	recordpetition := RecordPetition{
		Because:          "ต้องการเรียนจบไวขึ้น",
		RegisteredCredit: 24,
		TimeRecord:       time.Now().Add(24 * time.Hour), //ผิด
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(recordpetition)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลวันเวลาไม่ถูกต้อง"))
}
