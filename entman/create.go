package entman

import (
	"context"

	"github.com/yaroslav-koval/graphql-psychologists-courses/bunmodels"
	"github.com/yaroslav-koval/graphql-psychologists-courses/graph/model"
	"github.com/yaroslav-koval/graphql-psychologists-courses/pkg/logging"
)

func (em *EntityManager) CreatePsychologist(
	ctx context.Context, input model.NewPsychologist) (*bunmodels.Psychologist, error) {
	p := &bunmodels.Psychologist{
		Name:        input.Name,
		Description: input.Description,
	}

	_, err := em.db.NewInsert().
		Model(p).
		Exec(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return p, nil
}

func (em *EntityManager) CreateCourse(ctx context.Context, input model.NewCourse) (*bunmodels.Course, error) {
	c := &bunmodels.Course{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	tx, err := em.db.BeginTx(ctx, nil)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			logging.SendSimpleError(err)
		}
	}()

	_, err = tx.NewInsert().
		Model(c).
		Exec(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	relations := []*bunmodels.CoursePsychologist{}
	for _, pid := range input.Psychologists {
		relations = append(relations, &bunmodels.CoursePsychologist{
			CourseID:       c.ID,
			PsychologistID: pid,
		})
	}

	_, err = tx.NewInsert().
		Model(&relations).
		Exec(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return c, nil
}

func (em *EntityManager) CreateLesson(ctx context.Context, input model.NewLesson) (*bunmodels.Lesson, error) {
	l := &bunmodels.Lesson{
		Name:     input.Name,
		Number:   input.Number,
		CourseID: input.Course,
	}

	_, err := em.db.NewInsert().
		Model(l).
		Exec(ctx)
	if err != nil {
		logging.SendSimpleErrorAsync(err)
		return nil, err
	}

	return l, nil
}
