package app

import "lawrenceMichaelMargaja/newStudent/api/controllers"

func mapUrls()  {
	//Student
	router.GET("/student", controllers.StudentController.GetStudents)
	router.POST("/student", controllers.StudentController.Create)
	router.POST("/student/enroll", controllers.StudentController.Enroll)

	//Course
	router.GET("/course/:course_id", controllers.CoursesController.GetEnrolledStudents)
	router.POST("/course", controllers.CoursesController.Create)
	router.DELETE("/course/:course_id", controllers.CoursesController.DeleteCourse)
}
