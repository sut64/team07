package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ข้อมูลถูกต้องหมดทุก field
func TestPass(t *testing.T) {
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

	wd := Withdrawal{
		YearTime:       20400,
		RemainCredit:   16,
		Reason:         "so hard",
		WithdrawalTime: time.Now(),
	}
	ok, err := govalidator.ValidateStruct(wd)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ข้อมูลปีการศึกษาผิดพลาด"))
}

// ข้อมูลหน่วยกิตไม่ถูกต้อง ไม่เป็นจำนวนเต็มบวก
func TestRemainCreditIsPosiviteNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	wd := Withdrawal{
		YearTime:       2564,
		RemainCredit:   -5,
		Reason:         "so hard",
		WithdrawalTime: time.Now(),
	}
	ok, err := govalidator.ValidateStruct(wd)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ข้อมูลหน่วยกิตไม่ถูกต้อง"))
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

	wd := Withdrawal{
		YearTime:       2564,
		RemainCredit:   16,
		Reason:         "so hard",
		WithdrawalTime: time.Now().Add(24 * time.Hour),
	}

	ok, err := govalidator.ValidateStruct(wd)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ข้อมูลวันเวลาไม่ถูกต้อง"))
}
