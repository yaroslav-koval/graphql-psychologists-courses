package entman

import (
	"context"

	"github.com/yaroslav-koval/graphql-psychologists-courses/bunmodels"
	"github.com/yaroslav-koval/graphql-psychologists-courses/graph/model"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

func (em *EntityManager) UpdatePsychologist(
	ctx context.Context, input model.UpdatePsychologist) (*bunmodels.Psychologist, error) {
	p := &bunmodels.Psychologist{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}

	_, err := em.db.NewUpdate().
		Model(p).
		WherePK().
		OmitZero().
		Exec(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return p, nil
}

func (em *EntityManager) UpdateCourse(ctx context.Context, input model.UpdateCourse) (*bunmodels.Course, error) {
	c := &bunmodels.Course{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	_, err := em.db.NewUpdate().
		Model(c).
		WherePK().
		OmitZero().
		Exec(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return c, nil
}

func (em *EntityManager) UpdateLesson(ctx context.Context, input model.UpdateLesson) (*bunmodels.Lesson, error) {
	l := &bunmodels.Lesson{
		ID:     input.ID,
		Name:   input.Name,
		Number: input.Number,
	}

	_, err := em.db.NewUpdate().
		Model(l).
		WherePK().
		OmitZero().
		Returning("course").
		Exec(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return l, nil
}
