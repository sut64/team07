package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ข้อมูลถูกต้องหมดทุก field
func TestWithdrawalPass(t *testing.T) {
	g := NewGomegaWithT(t)

	withdrawal := Withdrawal{
		YearTime:       2564,
		RemainCredit:   16,
		Reason:         "work hard",
		WithdrawalTime: time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(withdrawal)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

// ข้อมูลปีการศึกษาผิดพลาด
func TestYearTimeIsNot4Digit(t *testing.T) {

	g := NewGomegaWithT(t)
	fixtures := []int{
		-2564,
		25555,
		255,
		25,
		0,
	}

	for _, fixture := range fixtures {

		withdrawal := Withdrawal{
			YearTime:       fixture,
			RemainCredit:   16,
			Reason:         "so hard",
			WithdrawalTime: time.Now(),
		}
		ok, err := govalidator.ValidateStruct(withdrawal)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("ข้อมูลปีการศึกษาผิดพลาด"))
	}
}

// ข้อมูลหน่วยกิตไม่ถูกต้อง
func TestRemainCreditIsNotPosiviteNumber(t *testing.T) {
	g := NewGomegaWithT(t)
	fixtures := []int{
		-20,
		0,
	}
	for _, fixture := range fixtures {
		withdrawal := Withdrawal{
			YearTime:       2564,
			RemainCredit:   fixture,
			Reason:         "so hard",
			WithdrawalTime: time.Now(),
		}
		ok, err := govalidator.ValidateStruct(withdrawal)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("ข้อมูลหน่วยกิตไม่ถูกต้อง"))
	}
}

// ข้อมูลเหตุผลไม่ถูกต้อง
func TestReasonNotBlank(t *testing.T) {

	g := NewGomegaWithT(t)

	withdrawal := Withdrawal{
		YearTime:       2564,
		RemainCredit:   16,
		Reason:         "",
		WithdrawalTime: time.Now(),
	}

	ok, err := govalidator.ValidateStruct(withdrawal)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ข้อมูลเหตุผลไม่ถูกต้อง"))

}

// ข้อมูลวันที่เวลาไม่ถูกต้อง
func TestWithdrawalTimeMustBePresent(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []time.Time{
		time.Now().Add(24 * time.Hour),
		time.Now().Add(24 - time.Hour),
	}

	for _, fixture := range fixtures {

		withdrawal := Withdrawal{
			YearTime:       2564,
			RemainCredit:   16,
			Reason:         "so hard",
			WithdrawalTime: fixture,
		}

		ok, err := govalidator.ValidateStruct(withdrawal)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("ข้อมูลวันเวลาไม่ถูกต้อง"))
	}
}
