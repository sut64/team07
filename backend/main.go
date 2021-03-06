package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sut64/team07/controller"
	"github.com/sut64/team07/entity"
	"github.com/sut64/team07/middlewares"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Student Routes
			protected.GET("/students", controller.ListStudents)
			protected.GET("/students/:id", controller.ListStudent)
			protected.GET("/student/:id", controller.GetStudent)
			protected.PATCH("/students", controller.UpdateStudent)
			protected.DELETE("/students/:id", controller.DeleteStudent)

			// Teaher Routes
			protected.GET("/teachers", controller.ListTeacher)
			protected.GET("/teacher/:id", controller.GetTeacher)
			protected.POST("/teacher", controller.CreateTeacher)
			protected.PATCH("/teachers", controller.UpdateTeacher)
			protected.DELETE("/teachers/:id", controller.DeleteTeacher)

			// Registrar Routes
			protected.GET("/registrars", controller.ListRegistrar)
			protected.GET("/registrar/:id", controller.GetRegistrar)
			protected.POST("/registrar", controller.CreateRegistrar)
			protected.PATCH("/registrars", controller.UpdateRegistrar)
			protected.DELETE("/registrars/:id", controller.DeleteRegistrar)

			// Semester Routes
			protected.GET("/semesters", controller.ListSemesters)
			protected.GET("/semester/:id", controller.GetSemester)
			protected.POST("/semesters", controller.CreateSemester)
			protected.PATCH("/semesters", controller.UpdateSemester)
			protected.DELETE("/semesters/:id", controller.DeleteSemester)

			// ExamType Routes
			protected.GET("/examtypes", controller.ListExamTypes)
			protected.GET("/examtype/:id", controller.GetExamType)
			protected.POST("/examtypes", controller.CreateExamType)
			protected.PATCH("/examtypes", controller.UpdateExamType)
			protected.DELETE("/examtypes/:id", controller.DeleteExamType)

			// ExamSchedule Routes
			protected.GET("/examschedules", controller.ListExamSchedules)
			protected.GET("/examschedule/:id", controller.GetExamSchedule)
			protected.POST("/examschedules", controller.CreateExamSchedule)
			protected.PATCH("/examschedules", controller.UpdateExamSchedule)
			protected.DELETE("/examschedules/:id", controller.DeleteExamSchedule)

			// Course Routes
			protected.GET("/courses", controller.ListCourses)
			protected.GET("/course/:id", controller.GetCourse)
			protected.POST("/courses", controller.CreateCourse)
			protected.PATCH("/courses", controller.UpdateCourse)
			protected.DELETE("/courses/:id", controller.DeleteCourse)

			// RegisCourse Routes
			protected.GET("/regiscourses", controller.ListRegisCoursess)
			protected.GET("/regiscourses/:id", controller.ListRegisCourses)
			protected.GET("/regiscourse/:id", controller.GetRegisCourse)
			protected.POST("/regiscourses", controller.CreateRegisCourse)
			protected.PATCH("/regiscourses", controller.UpdateRegisCourse)
			protected.DELETE("/regiscourses/:id", controller.DeleteRegisCourse)

			// AddCourse Routes
			protected.GET("/addcourses", controller.ListAddCourses)
			protected.GET("/addcourse/:id", controller.GetAddCourse)
			protected.POST("/addcourse", controller.CreateAddCourse)
			protected.PATCH("/addcourses", controller.UpdateAddCourse)
			protected.DELETE("/addcourses/:id", controller.DeleteAddCourse)

			// Program Routes
			protected.GET("/programs", controller.ListPrograms)
			protected.GET("/program/:id", controller.GetProgram)
			protected.POST("/programs", controller.CreateProgram)
			protected.PATCH("/programs", controller.UpdateProgram)
			protected.DELETE("/programs/:id", controller.DeleteProgram)

			// Withdrwals Routes
			protected.GET("/withdrawals", controller.ListWithdrawals)
			protected.GET("/withdrawals/:id", controller.ListWithdrawal)
			protected.GET("/withdrawal/:id", controller.GetWithdrwal)
			protected.POST("/withdrawal", controller.CreateWithdrawal)
			protected.PATCH("/withdrawals", controller.UpdateWithdrawal)
			protected.DELETE("/withdrawals/:id", controller.DeleteWithdrawal)

			// RequestStatus Routes
			protected.GET("/requeststatuses", controller.ListRequestStatuses)
			protected.GET("/requeststatus/:id", controller.GetRequestStatus)
			protected.POST("/requeststatuses", controller.CreateRequestStatus)
			protected.PATCH("/requeststatuses", controller.UpdateRequestStatus)
			protected.DELETE("/requeststatuses/:id", controller.DeleteRequestStatus)

			// RequestExam Routes
			protected.GET("/request_exams", controller.ListRequestExams)
			protected.GET("/requestexam/:id", controller.GetRequestExam)
			protected.POST("/request_exams", controller.CreateRequestExam)
			protected.PATCH("/request_exams", controller.UpdateRequestExam)
			protected.DELETE("/requestexams/:id", controller.DeleteRequestExam)

			// Petition Routes
			protected.GET("/petitions", controller.ListPetitions)
			protected.GET("/petition/:id", controller.GetPetition)
			protected.POST("/petitions", controller.CreatePetition)
			protected.PATCH("/petitions", controller.UpdatePetition)
			protected.DELETE("/petitions/:id", controller.DeletePetition)

			// RecordPetition Routes
			protected.GET("/record_petitions", controller.ListRecordPetitions)
			protected.GET("/recordpetition/:id", controller.GetRecordPetition)
			protected.POST("/record_petitions", controller.CreateRecordPetition)
			protected.PATCH("/record_petitions", controller.UpdateRecordPetition)
			protected.DELETE("/recordpetitions/:id", controller.DeleteRecordPetition)

			// Grades Routes
			protected.GET("/grades", controller.ListGrades)
			protected.GET("/grades/:id", controller.GetGrades)
			protected.POST("/grades", controller.CreateGrades)
			protected.PATCH("/grades", controller.UpdateGrades)
			protected.DELETE("/grades/:id", controller.DeleteGrades)

			// IncreaseGrades Routes
			protected.GET("/increasegrades", controller.ListIncreaseGrades)
			protected.GET("/increasegrade/:id", controller.GetIncreaseGrades)
			protected.POST("/increasegrades", controller.CreateIncreaseGrades)
			protected.PATCH("/increasegrades", controller.UpdateIncreaseGrades)
			protected.DELETE("/increasegrades/:id", controller.DeleteSemester)

		}
	}

	// Authentication Routes
	r.POST("/student/login", controller.LoginStudent)
	r.POST("/registrar/login", controller.LoginRegistrar)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}

}
