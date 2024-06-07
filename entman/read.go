package entman

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yaroslav-koval/graphql-psychologists-courses/bunmodels"
	"github.com/yaroslav-koval/graphql-psychologists-courses/graph/model"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

func (em *EntityManager) GetPsychologistByID(ctx context.Context, id string) (*model.Psychologist, error) {
	panic(fmt.Errorf("not implemented: Psychologist - psychologist"))
}

func (em *EntityManager) GetAllPsychologists(ctx context.Context) ([]*model.Psychologist, error) {
	panic(fmt.Errorf("not implemented: Psychologists - psychologists"))
}

func (em *EntityManager) GetCourseByID(ctx context.Context, id string) (*model.Course, error) {
	panic("1")
}

func (em *EntityManager) GetAllCourses(ctx context.Context) ([]*model.Course, error) {
	panic("1")
}

func (em *EntityManager) GetLessonByID(ctx context.Context, id string) (*model.Lesson, error) {
	l := &bunmodels.Lesson{}

	err := em.db.NewSelect().
		Model(l).
		Where("l.id = ?", id).
		Relation("Course").
		Relation("Course.Lessons").
		Relation("Course.Psychologists").
		Scan(ctx)
	if err != nil {
		logging.SendAsync(logging.Error().Str("error", err.Error()))
		return nil, err
	}

	res := &model.Lesson{}
	err = parseBunModelToGenerated(l, res)
	if err != nil {
		logging.SendAsync(logging.Error().Str("error", err.Error()))
		return nil, err
	}

	return res, nil
}

func (em *EntityManager) GetAllLessons(ctx context.Context) ([]*model.Lesson, error) {
	l := []*bunmodels.Lesson{}

	err := em.db.NewSelect().
		Model(&l).
		Relation("Course").
		Relation("Course.Lessons").
		Relation("Course.Psychologists").
		Scan(ctx)
	if err != nil {
		logging.SendAsync(logging.Error().Str("error", err.Error()))
		return nil, err
	}

	res := []*model.Lesson{}
	for _, item := range l {
		r := &model.Lesson{}
		err = parseBunModelToGenerated(item, r)
		if err != nil {
			logging.SendAsync(logging.Error().Str("error", err.Error()))
			return nil, err
		}

		res = append(res, r)
	}

	return res, nil
}

func parseBunModelToGenerated(source, destination any) error {
	m, err := json.Marshal(source)
	if err != nil {
		logging.SendAsync(logging.Error().Str("error", err.Error()))
		return err
	}

	err = json.Unmarshal(m, &destination)
	if err != nil {
		logging.SendAsync(logging.Error().Str("error", err.Error()))
		return err
	}

	return nil
}
