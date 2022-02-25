package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestExamSchedulePass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	examschedule := ExamSchedule{
		AcademicYear: 2564,
		ExamRoom: "B5204",
		ExamDate: time.Now().Add(time.Hour*24),
		StartTime:time.Date(2021, 1, 20, 11, 30, 0, 0, time.Local),
		EndTime: time.Date(2021, 1, 20, 12, 30, 0, 0, time.Local),
		
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(examschedule)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

//ตรวจสอบปีการศึกษาต้องตัวเลข 4 หลัก
func TestAcademicYearMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []int{
		-2564,
		1,
		12,
		123,
		12345,
		123456,
		0,
	}

	for _, fixture := range fixtures {
		examschedule := ExamSchedule{
			AcademicYear: fixture, //ผิด
			ExamRoom: "B5204", 
			ExamDate: time.Now().Add(time.Hour*24),
			StartTime:time.Date(2021, 1, 20, 11, 30, 0, 0, time.Local),
			EndTime: time.Date(2021, 1, 20, 12, 30, 0, 0, time.Local),
		}

		ok, err := govalidator.ValidateStruct(examschedule)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("ปีการศึกษาต้องเป็นตัวเลข 4 หลัก"))
	}
}	

// ตรวจสอบห้องสอบต้องขึ้นต้นด้วย B และตามด้วยตัวเลข 4 ตัว
func TestExamRoomMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"X60000",
		"BA0000",  // B ตามด้วย A และ \d 4 ตัว
		"B00000",   // B ตามด้วย \d 5 ตัว
		"B000000", // B ตามด้วย \d 6 ตัว
		"B0000000", // B ตามด้วย \d 7 ตัว
		"F11-422", // ขึ้นต้น F 
		"B", //ตัวอักษร 1 ตัว
		"11111", //มีแต่ตัวเลข
	}

	for _, fixture := range fixtures {
		examschedule := ExamSchedule{
			AcademicYear: 2564,
			ExamRoom: fixture, //ผิด
			ExamDate: time.Now().Add(time.Hour*24),
			StartTime:time.Date(2021, 1, 20, 11, 30, 0, 0, time.Local),
			EndTime: time.Date(2021, 1, 20, 12, 30, 0, 0, time.Local),
		}

		ok, err := govalidator.ValidateStruct(examschedule)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("ห้องสอบต้องขึ้นต้นด้วย B และตามด้วยตัวเลข 4 ตัว"))
	}
}

//ตรวจสอบวันที่สอบต้องเป็นวันที่ในอนาคต
func TestDateExamMustBeFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []time.Time{
		time.Now().Add(-24 * time.Hour),
		time.Now(), 
	}

	for _, fixture := range fixtures {
		examschedule := ExamSchedule{
			AcademicYear: 2564,
			ExamRoom: "B5204",
			ExamDate: fixture,
			StartTime:time.Date(2021, 1, 20, 11, 30, 0, 0, time.Local),
			EndTime: time.Date(2021, 1, 20, 12, 30, 0, 0, time.Local),
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(examschedule)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("วันที่สอบต้องเป็นวันในอนาคต"))
	}
}

