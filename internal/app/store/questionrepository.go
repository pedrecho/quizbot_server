package store

import "basehttpserver/internal/app/model"

type QuestionRepository struct {
	store *Store
}

func (r *QuestionRepository) Create(u *model.Question) (*model.Question, error) {
	query := "INSERT INTO questions (package, text, option1, option2, option3, option4, answer) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	if err := r.store.db.QueryRow(query,
		u.Package,
		u.Text,
		u.Option1,
		u.Option2,
		u.Option3,
		u.Option4,
		u.Answer,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *QuestionRepository) FindPackage(pack string) ([]*model.Question, error) {
	return nil, nil
}
