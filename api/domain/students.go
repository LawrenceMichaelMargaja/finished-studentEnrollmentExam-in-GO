package domain

type Student struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Email          string `json:"email" db:"email"`
	Phone          string `json:"phone" db:"phone"`
	CourseEnrolled string `json:"courses_enrolled" db:"courses_enrolled"`
}

type CreateStudent struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type EnrollStudent struct {
	StudentId int `json:"student_id"`
	CourseId  int `json:"course_id"`
}
