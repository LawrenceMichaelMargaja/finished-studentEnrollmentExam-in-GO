package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"lawrenceMichaelMargaja/newStudent/api/utils"
)


type studentDaoInterface interface {
	GetStudents() (*[]Student, error)
	Create(name string, email string, phone string) error
	IsValidId(studentId int) (bool, error)
	Enroll(studentId int, courseId int) error
	SetClient()
}

type studentDao struct {
	client *sqlx.DB
}

var (
	StudentDao studentDaoInterface
)

func init() {
	StudentDao = &studentDao{}
	StudentDao.SetClient()
}

func (s *studentDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

func (s *studentDao) GetStudents() (*[]Student, error) {
	var students []Student
	sql := `
		SELECT
		  s.student_id AS id,
		  s.student_name AS ` + utils.EncloseString("name", "`") + `,
		  s.student_email AS email,
		  s.student_phone AS phone,
		  IF(GROUP_CONCAT(c.course_name) IS NULL, "", GROUP_CONCAT(c.course_name)) AS courses_enrolled
		FROM
		  student s
		  LEFT JOIN students_enrolled se
			ON 1 = 1
			AND s.student_id = se.student_ref_id
		  LEFT JOIN course c
			ON 1 = 1
			AND se.course_ref_id = c.course_id
		GROUP BY s.student_id
	`

	err := s.client.Select(&students, sql)

	return &students, err
}

func (s *studentDao) Create(name string, email string, phone string) error {
	sql := `
		INSERT INTO student (
		  student_name,
		  student_email,
		  student_phone
		)
		VALUES
		  (
			?,
			?,
			?
		  );
	`
	_,err := s.client.Exec(sql, name, email, phone)
	return err
}

func (s *studentDao) Enroll(studentId int, courseId int) error {
	sql := `INSERT INTO students_enrolled(student_ref_id, course_ref_id) VALUES(?, ?);`
	_, err := s.client.Exec(sql, studentId, courseId)

	return err
}

func (s *studentDao) IsValidId(id int) (bool, error) {
	var count2 int
	sql := `
		SELECT COUNT(*) AS count2
		FROM student 
		WHERE student_id = ?
	`

	err := s.client.Get(&count2, sql, id)

	if err != nil {
		return false, err
	}

	if count2 == 0 {
		return false, nil
	}

	return true, nil
}