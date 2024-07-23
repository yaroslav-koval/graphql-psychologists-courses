package entman

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yaroslav-koval/graphql-psychologists-courses/graph/model"
)

func (s *entmanSuite) TestCreatePsychologist() {
	expID := "1"
	expCreatedAt := "created date"
	expUpdatedAt := "updated date"
	s.mock.ExpectQuery(`INSERT INTO "psychologist"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(expID, expCreatedAt, expUpdatedAt),
		)

	d := "description value"
	expPsycho := model.NewPsychologist{
		Name:        "name value",
		Description: &d,
	}
	p, err := s.em.CreatePsychologist(cb, expPsycho)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(expID, p.ID)
	s.Equal(expCreatedAt, p.CreatedAt)
	s.Equal(expUpdatedAt, p.UpdatedAt)
	s.Equal(expPsycho.Name, p.Name)
	s.EqualValues(expPsycho.Description, p.Description)
}

func (s *entmanSuite) TestCreateCourseSuccess() {
	s.mock.ExpectBegin()

	expID := "1"
	expCreatedAt := "created date"
	expUpdatedAt := "updated date"
	s.mock.ExpectQuery(`INSERT INTO "course"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(expID, expCreatedAt, expUpdatedAt),
		)

	s.mock.ExpectCommit()

	d := "description value"
	expC := model.NewCourse{
		Name:        "name value",
		Description: &d,
		Price:       50,
	}
	c, err := s.em.CreateCourse(cb, expC)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(expID, c.ID)
	s.Equal(expCreatedAt, c.CreatedAt)
	s.Equal(expUpdatedAt, c.UpdatedAt)
	s.Equal(expC.Name, c.Name)
	s.EqualValues(expC.Description, c.Description)
	s.Equal(expC.Price, c.Price)
}

func (s *entmanSuite) TestCreateCourseSuccessWithRelations() {
	s.mock.ExpectBegin()

	expID := "1"
	expCreatedAt := "created date"
	expUpdatedAt := "updated date"
	s.mock.ExpectQuery(`INSERT INTO "course"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(expID, expCreatedAt, expUpdatedAt),
		)

	expRelationID := "10"
	s.mock.ExpectQuery(`INSERT INTO "courses_psychologists"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(expRelationID, "", ""),
		)
	s.mock.ExpectCommit()

	d := "description value"
	expC := model.NewCourse{
		Name:          "name value",
		Description:   &d,
		Price:         50,
		Psychologists: []string{"1"},
	}
	c, err := s.em.CreateCourse(cb, expC)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(expID, c.ID)
	s.Equal(expCreatedAt, c.CreatedAt)
	s.Equal(expUpdatedAt, c.UpdatedAt)
	s.Equal(expC.Name, c.Name)
	s.EqualValues(expC.Description, c.Description)
	s.Equal(expC.Price, c.Price)
}

func (s *entmanSuite) TestCreateCourseRollback() {
	s.mock.ExpectBegin()

	expErr := fmt.Errorf("expected error")
	s.mock.ExpectQuery(`INSERT INTO "course"`).
		WillReturnError(expErr)

	s.mock.ExpectRollback()

	d := "description value"
	expC := model.NewCourse{
		Name:        "name value",
		Description: &d,
		Price:       50,
	}
	_, err := s.em.CreateCourse(cb, expC)
	s.ErrorIs(expErr, err)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *entmanSuite) TestCreateLesson() {
	expID := "1"
	expCreatedAt := "created date"
	expUpdatedAt := "updated date"
	s.mock.ExpectQuery(`INSERT INTO "lesson"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(expID, expCreatedAt, expUpdatedAt),
		)

	expLesson := model.NewLesson{
		Name:   "name value",
		Number: 2,
		Course: "course_id_123",
	}
	l, err := s.em.CreateLesson(cb, expLesson)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(expID, l.ID)
	s.Equal(expCreatedAt, l.CreatedAt)
	s.Equal(expUpdatedAt, l.UpdatedAt)
	s.Equal(expLesson.Name, l.Name)
	s.Equal(expLesson.Number, l.Number)
	s.Equal(expLesson.Course, l.CourseID)
}
