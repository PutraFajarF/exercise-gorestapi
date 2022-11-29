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

type CreateQuestionInput struct {
	Body          string `json:"body"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	CorrectAnswer string `json:"correct_answer"`
	Score         int    `json:"score"`
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

func NewQuestion(exerciseID, score, creatorID int, body, a, b, c, d, correctAnswer string) (*Question, error) {
	if body == "" {
		return nil, errors.New("body is required")
	}

	if a == "" {
		return nil, errors.New("option a is required")
	}

	if b == "" {
		return nil, errors.New("option b is required")
	}

	if c == "" {
		return nil, errors.New("option c is required")
	}

	if d == "" {
		return nil, errors.New("option d is required")
	}

	if correctAnswer == "" {
		return nil, errors.New("correct answer is required")
	}

	if score == 0 {
		return nil, errors.New("score must be more than 0")
	}

	question := &Question{
		ExerciseID:    exerciseID,
		Body:          body,
		OptionA:       a,
		OptionB:       b,
		OptionC:       c,
		OptionD:       d,
		CorrectAnswer: correctAnswer,
		Score:         score,
		CreatorID:     creatorID,
	}

	return question, nil
}
