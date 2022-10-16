package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config             *Config
	db                 *sql.DB
	questionRepository *QuestionRepository
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DataBaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Question() *QuestionRepository {
	if s.questionRepository != nil {
		return s.questionRepository
	}
	s.questionRepository = &QuestionRepository{
		store: s,
	}
	return s.questionRepository
}
