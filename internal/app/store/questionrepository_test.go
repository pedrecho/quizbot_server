package store

import (
	"basehttpserver/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuestionRepository_Create(t *testing.T) {
	s, teardown := TestStore(t, databaseURL)
	defer teardown("questions")
	u, err := s.Question().Create(&model.Question{
		Package: "first",
		Text:    "Hello?",
		Option1: "YES",
		Option2: "YES",
		Option3: "YES",
		Option4: "YES",
		Answer:  1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
