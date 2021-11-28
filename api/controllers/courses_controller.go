package controllers

import (
	"github.com/gin-gonic/gin"
	"lawrenceMichaelMargaja/newStudent/api/domain"
	"lawrenceMichaelMargaja/newStudent/api/services"
	"lawrenceMichaelMargaja/newStudent/api/utils"
	"net/http"
	"strconv"
)

type coursesController struct {
}

var (
	CoursesController *coursesController
)

func init() {
	CoursesController = &coursesController{}
}

func (controller *coursesController) GetEnrolledStudents(c *gin.Context) {
	courseId, err := strconv.ParseInt(c.Param("course_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "course_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	students, err := services.CourseService.GetEnrolledStudents(int(courseId))

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to fetch the enrolled students : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, students)
	return
}

func (controller *coursesController) Create(c *gin.Context) {
	var body domain.CreateCourse

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err := services.CourseService.Create(body.Name, body.Professor, body.Description)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to insert course data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, "Successfully created the course!")
	return
}

func (controller *coursesController) DeleteCourse(c *gin.Context) {
	courseId, err := strconv.ParseInt(c.Param("course_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "course_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	err = services.CourseService.DeleteCourse(int(courseId))

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to delete the course : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, "Successfully deleted the course!")
}
