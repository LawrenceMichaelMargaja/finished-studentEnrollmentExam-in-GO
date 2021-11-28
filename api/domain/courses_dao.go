package domain

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"lawrenceMichaelMargaja/newStudent/api/utils"
)

type courseDaoInterface interface {
	GetEnrolledStudents(courseId int) (*[]Student, error)
	Create(name string, professor string, description string) error
	DeleteCourse(courseId int) error
	SetClient()
	HasStudentEnrolled(courseId int) (bool, error)
	IsValidId(courseId int) (bool, error)
}

type courseDao struct {
	client *sqlx.DB
}

var (
	CourseDao courseDaoInterface
)

func init() {
	CourseDao = &courseDao{}
	CourseDao.SetClient()
}

func (s *courseDao) SetClient()  {
	s.client = utils.GetMYSQLConnection()
}

func (s *courseDao) GetEnrolledStudents(courseId int) (*[]Student, error) {
	var students []Student
	sql := `SELECT s.student_id AS id, s.student_name AS name, s.student_email AS email, s.student_phone AS phone,
			IF(GROUP_CONCAT(c.course_name) IS NULL, "", GROUP_CONCAT(c.course_name)) AS courses_enrolled
			FROM student s
			LEFT JOIN students_enrolled se ON 1 = 1 AND s.student_id = se.student_ref_id
			LEFT JOIN course c ON 1 = 1 AND se.course_ref_id = c.course_id
			WHERE 1 = 1 AND se.course_ref_id = ?
			GROUP BY s.student_id`

	err := s.client.Select(&students, sql, courseId)

	if len(students) == 0 {
		return nil, errors.New("no students enrolled in this course_id")
	}

	return &students, err
}

func (s *courseDao) DeleteCourse(id int) error {
	sql := `DELETE FROM course WHERE course_id = ?`

	_, err := s.client.Exec(sql, id)
	return err
}

func (s *courseDao) Create(name string, professor string, description string) error {
	sql := `INSERT INTO course(course_name, course_professor, course_description) VALUES(?, ?, ?);`
	_, err := s.client.Exec(sql, name, professor, description)
	return err
}

func (s *courseDao) IsValidId(id int) (bool, error) {
	var count int
	sql := `SELECT COUNT(*) AS count FROM course WHERE course_id = ?`

	err := s.client.Get(&count, sql, id)

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (s *courseDao) HasStudentEnrolled(id int) (bool, error) {
	var count int
	sql := `SELECT COUNT(*) AS count FROM students_enrolled WHERE course_ref_id = ?`

	err := s.client.Get(&count, sql, id)

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}