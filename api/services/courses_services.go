package services

import (
	"errors"
	"lawrenceMichaelMargaja/newStudent/api/domain"
)

type courseService struct {
	
}

var (
	CourseService *courseService
)

func init()  {
	CourseService = &courseService{}
}

func (s *courseService) GetEnrolledStudents(courseId int) (*[]domain.Student, error) {
	if courseId == 0 {
		return nil, errors.New("student_id is invalid")
	}

	isValidCourseId, err := domain.CourseDao.IsValidId(courseId)

	if err != nil {
		return nil, err
	}

	if isValidCourseId == false {
		return nil, errors.New("course_id is invalid")
	}

	return domain.CourseDao.GetEnrolledStudents(courseId)
}

func (s *courseService) Create(name string, professor string, description string) error {
	if name == "" {
		return errors.New("name is invalid")
	}
	if professor == "" {
		return errors.New("professor is invalid")
	}
	if description == "" {
		return errors.New("description is invalid")
	}

	return domain.CourseDao.Create(name, professor, description)
}

func (s *courseService) DeleteCourse(courseId int) error {
	if courseId == 0 {
		return errors.New("course_id is invalid")
	}

	isValidCourseId, err := domain.CourseDao.IsValidId(courseId)

	if err != nil {
		return err
	}

	if isValidCourseId == false {
		return errors.New("course_id is invalid")
	}

	hasStudentsEnrolled, err := domain.CourseDao.HasStudentEnrolled(courseId)

	if err != nil {
		return err
	}

	if hasStudentsEnrolled == true {
		return errors.New("cannot delete a course with students enrolled")
	}

	return domain.CourseDao.DeleteCourse(courseId)
}