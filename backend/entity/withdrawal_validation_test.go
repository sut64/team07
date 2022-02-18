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
	ok, err := govalidator.ValidateStruct(withdrawal)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}

// ข้อมูลปีการศึกษาผิดพลาด ไม่เป็นตัวเลข 4 หลัก
func TestYearTimeIsNot4Digit(t *testing.T) {

	g := NewGomegaWithT(t)
	fixtures := []int{
		-2564,
		3000,
		25555,
		255,
		0,
	}

	for _, fixture := range fixtures {

		wd := Withdrawal{
			YearTime:       fixture,
			RemainCredit:   16,
			Reason:         "so hard",
			WithdrawalTime: time.Now(),
		}
		ok, err := govalidator.ValidateStruct(wd)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("ข้อมูลปีการศึกษาผิดพลาด"))
	}
}

// ข้อมูลหน่วยกิตไม่ถูกต้อง ไม่เป็นจำนวนเต็มบวก
func TestRemainCreditIsNotPosiviteNumber(t *testing.T) {
	g := NewGomegaWithT(t)
	fixtures := []int{
		-20,
		0,
	}
	for _, fixture := range fixtures {
		wd := Withdrawal{
			YearTime:       2564,
			RemainCredit:   fixture,
			Reason:         "so hard",
			WithdrawalTime: time.Now(),
		}
		ok, err := govalidator.ValidateStruct(wd)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("ข้อมูลหน่วยกิตไม่ถูกต้อง"))
	}
}

// ข้อมูลเหตุผลไม่ถูกต้อง ห้ามเป็นช่องว่าง
func TestReasonNotBlank(t *testing.T) {

	g := NewGomegaWithT(t)

	withdrawal := Withdrawal{
		YearTime:       2564,
		RemainCredit:   16,
		Reason:         "",
		WithdrawalTime: time.Now(),
	}

	ok, err := govalidator.ValidateStruct(withdrawal)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ข้อมูลเหตุผลไม่ถูกต้อง"))

}

// ข้อมูลวันที่เวลาไม่ถูกต้อง ไม่เป็นวันเวลาปัจจุบัน
func TestWithdrawalTimeMustBePresent(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []time.Time{
		time.Now().Add(24 * time.Hour),
		time.Now().Add(24 - time.Hour),
	}

	for _, fixture := range fixtures {

		wd := Withdrawal{
			YearTime:       2564,
			RemainCredit:   16,
			Reason:         "so hard",
			WithdrawalTime: fixture,
		}

		ok, err := govalidator.ValidateStruct(wd)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("ข้อมูลวันเวลาไม่ถูกต้อง"))
	}
}
