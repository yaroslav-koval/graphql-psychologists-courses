package entman

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yaroslav-koval/graphql-psychologists-courses/graph/model"
)

func (s *entmanSuite) TestUpdatePsychologist() {
	d := "description"
	expValue := model.UpdatePsychologist{
		ID:          "psychologist_id",
		Name:        "random name",
		Description: &d,
	}

	s.mock.ExpectExec(`UPDATE "psychologist".*"name" = 'random name', "description" = 'description'.*"id" = 'psychologist_id'`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	p, err := s.em.UpdatePsychologist(cb, expValue)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.EqualValues(expValue.ID, p.ID)
	s.EqualValues(expValue.Name, p.Name)
	s.EqualValues(expValue.Description, p.Description)
}

func (s *entmanSuite) TestUpdateCourse() {
	d := "description"
	expValue := model.UpdateCourse{
		ID:          "course_id",
		Name:        "random name",
		Description: &d,
		Price:       3000,
	}

	s.mock.ExpectExec(`UPDATE "course".*"name" = 'random name', "description" = 'description', "price" = 3000.*"id" = 'course_id'`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	p, err := s.em.UpdateCourse(cb, expValue)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.EqualValues(expValue.ID, p.ID)
	s.EqualValues(expValue.Name, p.Name)
	s.EqualValues(expValue.Description, p.Description)
	s.EqualValues(expValue.Price, p.Price)
}

func (s *entmanSuite) TestUpdateLesson() {
	expValue := model.UpdateLesson{
		ID:     "lesson_id",
		Name:   "random name",
		Number: 50,
	}

	expCourseID := "course_id"
	s.mock.ExpectQuery(`UPDATE "lesson".*"name" = 'random name', "number" = 50.*"id" = 'lesson_id'`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"course",
			}).AddRow(
				expCourseID,
			),
		)

	p, err := s.em.UpdateLesson(cb, expValue)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.EqualValues(expValue.ID, p.ID)
	s.EqualValues(expValue.Name, p.Name)
	s.EqualValues(expValue.Number, p.Number)
	s.EqualValues(expCourseID, p.CourseID)
}
