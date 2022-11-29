package domain

import (
	"errors"
	"time"
)

type Exercise struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"question"`
}

type CreateExerciseInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Question struct {
	ID            int       `json:"id"`
	ExerciseID    int       `json:"exercise_id"`
	Body          string    `json:"body"`
	OptionA       string    `json:"option_a"`
	OptionB       string    `json:"option_b"`
	OptionC       string    `json:"option_c"`
	OptionD       string    `json:"option_d"`
	CorrectAnswer string    `json:"correct_answer"`
	Score         int       `json:"score"`
	CreatorID     int       `json:"creator_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Answer struct {
	ID         int       `json:"id"`
	ExerciseID int       `json:"exercise_id"`
	QuestionID int       `json:"question_id"`
	UserID     int       `json:"user_id"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewExercise(title, description string) (*Exercise, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	if description == "" {
		return nil, errors.New("description is required")
	}

	exercise := &Exercise{
		Title:       title,
		Description: description,
	}

	return exercise, nil
}
