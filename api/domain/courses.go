package domain

type Course struct {
	Id          string `json:"id" db:"course_id""`
	Name        string `json:"name" db:"course_name"`
	Professor   string `json:"professor" db:"course_professor"`
	Description string `json:"description" db:"course_description""`
}

type CreateCourse struct {
	Name        string `json:"name" db:"name""`
	Professor   string `json:"professor" db:"professor"`
	Description string `json:"description" db:"description"`
}
