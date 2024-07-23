package entman

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yaroslav-koval/graphql-psychologists-courses/bunmodels"
)

func (s *entmanSuite) TestGetPsychologistByID() {
	d := "description"
	expValue := &bunmodels.Psychologist{
		ID:          "psychologist_id",
		CreatedAt:   "created date",
		UpdatedAt:   "updated date",
		Name:        "random name",
		Description: &d,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "psychologist".*"id" = 'psychologist_id'`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"description",
			}).AddRow(
				expValue.ID,
				expValue.CreatedAt,
				expValue.UpdatedAt,
				expValue.Name,
				expValue.Description,
			),
		)

	p, err := s.em.GetPsychologistByID(cb, expValue.ID)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.EqualValues(expValue, p)
}

func (s *entmanSuite) TestGetAllPsychologists() {
	d := "description"
	expValue1 := &bunmodels.Psychologist{
		ID:          "psychologist_id_1",
		CreatedAt:   "created date",
		UpdatedAt:   "updated date",
		Name:        "random name",
		Description: &d,
	}

	d2 := "description 2"
	expValue2 := &bunmodels.Psychologist{
		ID:          "psychologist_id_2",
		CreatedAt:   "created date 2",
		UpdatedAt:   "updated date 2",
		Name:        "random name 2",
		Description: &d2,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "psychologist"`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"description",
			}).AddRow(
				expValue1.ID,
				expValue1.CreatedAt,
				expValue1.UpdatedAt,
				expValue1.Name,
				expValue1.Description,
			).AddRow(
				expValue2.ID,
				expValue2.CreatedAt,
				expValue2.UpdatedAt,
				expValue2.Name,
				expValue2.Description),
		)

	ps, err := s.em.GetAllPsychologists(cb)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(2, len(ps))
	s.EqualValues(expValue1, ps[0])
	s.EqualValues(expValue2, ps[1])
}

func (s *entmanSuite) TestGetCourseByID() {
	d := "description"
	expValue := &bunmodels.Course{
		ID:          "course_id",
		CreatedAt:   "created date",
		UpdatedAt:   "updated date",
		Name:        "random name",
		Description: &d,
		Price:       3000,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "course".*"id" = 'course_id'`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"description",
				"price",
			}).AddRow(
				expValue.ID,
				expValue.CreatedAt,
				expValue.UpdatedAt,
				expValue.Name,
				expValue.Description,
				expValue.Price,
			),
		)

	p, err := s.em.GetCourseByID(cb, expValue.ID)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.EqualValues(expValue, p)
}

func (s *entmanSuite) TestGetAllCourses() {
	d := "description"
	expValue1 := &bunmodels.Course{
		ID:          "course_id_1",
		CreatedAt:   "created date",
		UpdatedAt:   "updated date",
		Name:        "random name",
		Description: &d,
		Price:       3000,
	}

	d2 := "description 2"
	expValue2 := &bunmodels.Course{
		ID:          "course_id_2",
		CreatedAt:   "created date 2",
		UpdatedAt:   "updated date 2",
		Name:        "random name 2",
		Description: &d2,
		Price:       4000,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "course"`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"description",
				"price",
			}).AddRow(
				expValue1.ID,
				expValue1.CreatedAt,
				expValue1.UpdatedAt,
				expValue1.Name,
				expValue1.Description,
				expValue1.Price,
			).AddRow(
				expValue2.ID,
				expValue2.CreatedAt,
				expValue2.UpdatedAt,
				expValue2.Name,
				expValue2.Description,
				expValue2.Price,
			),
		)

	cs, err := s.em.GetAllCourses(cb)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(2, len(cs))
	s.EqualValues(expValue1, cs[0])
	s.EqualValues(expValue2, cs[1])
}

func (s *entmanSuite) TestGetLessonByID() {
	expValue := &bunmodels.Lesson{
		ID:        "lesson_id",
		CreatedAt: "created date",
		UpdatedAt: "updated date",
		Name:      "lesson name",
		Number:    7,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "lesson".*"id" = 'lesson_id'`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"number",
			}).AddRow(
				expValue.ID,
				expValue.CreatedAt,
				expValue.UpdatedAt,
				expValue.Name,
				expValue.Number,
			),
		)

	p, err := s.em.GetLessonByID(cb, expValue.ID)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.EqualValues(expValue, p)
}

func (s *entmanSuite) TestGetAllLessons() {
	expValue1 := &bunmodels.Lesson{
		ID:        "lesson_id",
		CreatedAt: "created date",
		UpdatedAt: "updated date",
		Name:      "lesson name",
		Number:    7,
	}

	expValue2 := &bunmodels.Lesson{
		ID:        "lesson_id",
		CreatedAt: "created date",
		UpdatedAt: "updated date",
		Name:      "lesson name",
		Number:    8,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "lesson"`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"number",
			}).AddRow(
				expValue1.ID,
				expValue1.CreatedAt,
				expValue1.UpdatedAt,
				expValue1.Name,
				expValue1.Number,
			).AddRow(
				expValue2.ID,
				expValue2.CreatedAt,
				expValue2.UpdatedAt,
				expValue2.Name,
				expValue2.Number,
			),
		)

	cs, err := s.em.GetAllLessons(cb)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(2, len(cs))
	s.EqualValues(expValue1, cs[0])
	s.EqualValues(expValue2, cs[1])
}

func (s *entmanSuite) TestGetPsychologistCourses() {
	d := "description"
	psycho := &bunmodels.Psychologist{
		ID:          "psychologist_id",
		CreatedAt:   "created date",
		UpdatedAt:   "updated date",
		Name:        "random name",
		Description: &d,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "psychologist".*"id" = 'psychologist_id'`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"description",
			}).AddRow(
				psycho.ID,
				psycho.CreatedAt,
				psycho.UpdatedAt,
				psycho.Name,
				psycho.Description,
			),
		)

	d1 := "description"
	expCourse1 := &bunmodels.Course{
		ID:          "course_id_1",
		CreatedAt:   "created date",
		UpdatedAt:   "updated date",
		Name:        "random name",
		Description: &d1,
		Price:       3000,
	}

	d2 := "description 2"
	expCourse2 := &bunmodels.Course{
		ID:          "course_id_2",
		CreatedAt:   "created date 2",
		UpdatedAt:   "updated date 2",
		Name:        "random name 2",
		Description: &d2,
		Price:       4000,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "course".*IN \('psychologist_id'\)`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"psychologist",
				"course",
				"id",
				"created_at",
				"updated_at",
				"name",
				"description",
				"price",
			}).AddRow(
				psycho.ID,
				expCourse1.ID,
				expCourse1.ID,
				expCourse1.CreatedAt,
				expCourse1.UpdatedAt,
				expCourse1.Name,
				expCourse1.Description,
				expCourse1.Price,
			).AddRow(
				psycho.ID,
				expCourse2.ID,
				expCourse2.ID,
				expCourse2.CreatedAt,
				expCourse2.UpdatedAt,
				expCourse2.Name,
				expCourse2.Description,
				expCourse2.Price,
			),
		)

	cs, err := s.em.GetPsychologistCourses(cb, psycho.ID)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(2, len(cs))
	s.EqualValues(expCourse1, cs[0])
	s.EqualValues(expCourse2, cs[1])
}

func (s *entmanSuite) TestGetCourseLessons() {
	d := "description"
	course := &bunmodels.Course{
		ID:          "course_id",
		CreatedAt:   "created date",
		UpdatedAt:   "updated date",
		Name:        "random name",
		Description: &d,
		Price:       2000,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "course".*"id" = 'course_id'`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"description",
				"price",
			}).AddRow(
				course.ID,
				course.CreatedAt,
				course.UpdatedAt,
				course.Name,
				course.Description,
				course.Price,
			),
		)

	expLesson1 := &bunmodels.Lesson{
		ID:        "course_id_1",
		CreatedAt: "created date",
		UpdatedAt: "updated date",
		Name:      "random name",
		Number:    10,
		CourseID:  course.ID,
	}

	expLesson2 := &bunmodels.Lesson{
		ID:        "course_id_2",
		CreatedAt: "created date 2",
		UpdatedAt: "updated date 2",
		Name:      "random name 2",
		Number:    11,
		CourseID:  course.ID,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "lesson".*IN \('course_id'\)`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"number",
				"course",
			}).AddRow(
				expLesson1.ID,
				expLesson1.CreatedAt,
				expLesson1.UpdatedAt,
				expLesson1.Name,
				expLesson1.Number,
				course.ID,
			).AddRow(
				expLesson2.ID,
				expLesson2.CreatedAt,
				expLesson2.UpdatedAt,
				expLesson2.Name,
				expLesson2.Number,
				course.ID,
			),
		)

	cs, err := s.em.GetCourseLessons(cb, course.ID)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(2, len(cs))
	s.EqualValues(expLesson1, cs[0])
	s.EqualValues(expLesson2, cs[1])
}

func (s *entmanSuite) TestGetCoursePsychologists() {
	d := "description"
	course := &bunmodels.Course{
		ID:          "course_id",
		CreatedAt:   "created date",
		UpdatedAt:   "updated date",
		Name:        "random name",
		Description: &d,
		Price:       2000,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "course".*"id" = 'course_id'`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"created_at",
				"updated_at",
				"name",
				"description",
				"price",
			}).AddRow(
				course.ID,
				course.CreatedAt,
				course.UpdatedAt,
				course.Name,
				course.Description,
				course.Price,
			),
		)

	d1 := "description"
	expPsycho1 := &bunmodels.Psychologist{
		ID:          "psychologist_id_1",
		CreatedAt:   "created date",
		UpdatedAt:   "updated date",
		Name:        "random name",
		Description: &d1,
	}

	d2 := "description 2"
	expPsycho2 := &bunmodels.Psychologist{
		ID:          "psychologist_id_2",
		CreatedAt:   "created date 2",
		UpdatedAt:   "updated date 2",
		Name:        "random name 2",
		Description: &d2,
	}

	s.mock.ExpectQuery(`SELECT .* FROM "psychologist".*IN \('course_id'\)`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"course",
				"psychologist",
				"id",
				"created_at",
				"updated_at",
				"name",
				"description",
			}).AddRow(
				course.ID,
				expPsycho1.ID,
				expPsycho1.ID,
				expPsycho1.CreatedAt,
				expPsycho1.UpdatedAt,
				expPsycho1.Name,
				expPsycho1.Description,
			).AddRow(
				course.ID,
				expPsycho2.ID,
				expPsycho2.ID,
				expPsycho2.CreatedAt,
				expPsycho2.UpdatedAt,
				expPsycho2.Name,
				expPsycho2.Description,
			),
		)

	cs, err := s.em.GetCoursePsychologists(cb, course.ID)
	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
	s.Equal(2, len(cs))
	s.EqualValues(expPsycho1, cs[0])
	s.EqualValues(expPsycho2, cs[1])
}
