package entman

import (
	"context"

	"github.com/yaroslav-koval/graphql-psychologists-courses/bunmodels"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

func (em *EntityManager) GetPsychologistByID(ctx context.Context, id string) (*bunmodels.Psychologist, error) {
	p := &bunmodels.Psychologist{
		ID: id,
	}

	err := em.db.NewSelect().
		Model(p).
		WherePK().
		Scan(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return p, nil
}

func (em *EntityManager) GetAllPsychologists(ctx context.Context) ([]*bunmodels.Psychologist, error) {
	ps := []*bunmodels.Psychologist{}

	err := em.db.NewSelect().
		Model(&ps).
		Scan(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return ps, nil
}

func (em *EntityManager) GetCourseByID(ctx context.Context, id string) (*bunmodels.Course, error) {
	c := &bunmodels.Course{
		ID: id,
	}

	err := em.db.NewSelect().
		Model(c).
		WherePK().
		Scan(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return c, nil
}

func (em *EntityManager) GetAllCourses(ctx context.Context) ([]*bunmodels.Course, error) {
	cs := []*bunmodels.Course{}

	err := em.db.NewSelect().
		Model(&cs).
		Scan(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return cs, nil
}

func (em *EntityManager) GetLessonByID(ctx context.Context, id string) (*bunmodels.Lesson, error) {
	l := &bunmodels.Lesson{
		ID: id,
	}

	err := em.db.NewSelect().
		Model(l).
		WherePK().
		Scan(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return l, nil
}

func (em *EntityManager) GetAllLessons(ctx context.Context) ([]*bunmodels.Lesson, error) {
	ls := []*bunmodels.Lesson{}

	err := em.db.NewSelect().
		Model(&ls).
		Scan(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return ls, nil
}

func (em *EntityManager) GetPsychologistCourses(ctx context.Context, psychologistID string) ([]*bunmodels.Course, error) {
	p := &bunmodels.Psychologist{
		ID: psychologistID,
	}

	err := em.db.NewSelect().
		Model(p).
		WherePK().
		Relation("Courses").
		Scan(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return p.Courses, nil
}

func (em *EntityManager) GetCourseLessons(ctx context.Context, courseID string) ([]*bunmodels.Lesson, error) {
	c := &bunmodels.Course{
		ID: courseID,
	}

	err := em.db.NewSelect().
		Model(c).
		WherePK().
		Relation("Lessons").
		Scan(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return c.Lessons, nil
}

func (em *EntityManager) GetCoursePsychologists(ctx context.Context, courseID string) ([]*bunmodels.Psychologist, error) {
	c := &bunmodels.Course{
		ID: courseID,
	}

	err := em.db.NewSelect().
		Model(c).
		WherePK().
		Relation("Psychologists").
		Scan(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return c.Psychologists, nil
}
