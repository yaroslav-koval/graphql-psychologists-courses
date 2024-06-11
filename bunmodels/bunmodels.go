package bunmodels

import (
	"github.com/uptrace/bun"
)

func RegisterBunManyToManyModels(b *bun.DB) {
	b.RegisterModel((*CoursePsychologist)(nil))
}

type Psychologist struct {
	bun.BaseModel `bun:"table:psychologist,alias:p"`

	ID          string    `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`
	CreatedAt   string    `bun:",nullzero,default:current_timestamp" json:"created_at"`
	UpdatedAt   string    `bun:",nullzero,default:current_timestamp" json:"updated_at"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Courses     []*Course `bun:"m2m:courses_psychologists,join:Psychologist=Course" json:"courses"`
}

type Course struct {
	bun.BaseModel `bun:"table:course,alias:c"`

	ID            string          `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`
	CreatedAt     string          `bun:",nullzero,default:current_timestamp" json:"created_at"`
	UpdatedAt     string          `bun:",nullzero,default:current_timestamp" json:"updated_at"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Price         int             `json:"price"`
	Psychologists []*Psychologist `bun:"m2m:courses_psychologists,join:Course=Psychologist" json:"psychologists"`
	Lessons       []*Lesson       `bun:"rel:has-many,join:id=course" json:"lessons"`
}

type CoursePsychologist struct {
	bun.BaseModel `bun:"table:courses_psychologists,alias:cp"`

	ID             string        `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`
	CreatedAt      string        `bun:",nullzero,default:current_timestamp" json:"created_at"`
	UpdatedAt      string        `bun:",nullzero,default:current_timestamp" json:"updated_at"`
	CourseID       string        `bun:"course,type:uuid" json:"course_id"`
	Course         *Course       `bun:"rel:belongs-to,join:course=id" json:"course"`
	PsychologistID string        `bun:"psychologist,type:uuid" json:"psychologist_id"`
	Psychologist   *Psychologist `bun:"rel:belongs-to,join:psychologist=id" json:"psychologist"`
}

type Lesson struct {
	bun.BaseModel `bun:"table:lesson,alias:l"`

	ID        string  `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`
	CreatedAt string  `bun:",nullzero,default:current_timestamp" json:"created_at"`
	UpdatedAt string  `bun:",nullzero,default:current_timestamp" json:"updated_at"`
	Name      string  `json:"name"`
	Number    int     `json:"number"`
	CourseID  string  `bun:"course,type:uuid" json:"course_id"`
	Course    *Course `bun:"rel:belongs-to,join:course=id" json:"course"`
}
