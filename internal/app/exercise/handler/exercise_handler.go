package handler

import (
	"exercise/internal/app/domain"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseHandler struct {
	db *gorm.DB
}

func NewExerciseHandler(db *gorm.DB) *ExerciseHandler {
	return &ExerciseHandler{db: db}
}

func (eh ExerciseHandler) CreateExercise(c *gin.Context) {
	var newExercise domain.CreateExerciseInput

	if err := c.ShouldBind(&newExercise); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}

	exercise, err := domain.NewExercise(newExercise.Title, newExercise.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	if err := eh.db.Create(exercise).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Success created exercise",
	})
}

func (eh ExerciseHandler) GetExerciseByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
		return
	}

	var exercise domain.Exercise
	err = eh.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "exercise not found",
		})
		return
	}
	c.JSON(http.StatusOK, exercise)
}

func (eh ExerciseHandler) GetScore(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid id",
		})
		return
	}

	var exercise domain.Exercise
	err = eh.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "exercise not found",
		})
		return
	}
	userID := c.Request.Context().Value("user_id").(int)

	var answers []domain.Answer
	err = eh.db.Where("exercise_id = ? AND user_id = ?", id, userID).Find(&answers).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "not answere yet",
		})
		return
	}

	mapQA := make(map[int]domain.Answer)
	for _, answer := range answers {
		mapQA[answer.QuestionID] = answer
	}

	var score Score
	wg := new(sync.WaitGroup)
	for _, question := range exercise.Questions {
		wg.Add(1)
		go func(question domain.Question) {
			defer wg.Done()
			if strings.EqualFold(question.CorrectAnswer, mapQA[question.ID].Answer) {
				score.Inc(question.Score)
			}
		}(question)
	}

	wg.Wait()

	c.JSON(http.StatusOK, map[string]int{
		"score": score.total,
	})
}

type Score struct {
	total int
	mu    sync.Mutex
}

func (s *Score) Inc(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.total += value
}

func (eh ExerciseHandler) CreateQuestion(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid exercise id",
		})
		return
	}

	var exercise domain.Exercise
	err = eh.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "exercise data not found",
		})
		return
	}

	var inputQuestion domain.CreateQuestionInput
	if err := c.ShouldBind(&inputQuestion); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid body request",
		})
	}

	creatorID := c.Request.Context().Value("user_id").(int)

	question, err := domain.NewQuestion(id, inputQuestion.Score, creatorID, inputQuestion.Body, inputQuestion.OptionA, inputQuestion.OptionB, inputQuestion.OptionC, inputQuestion.OptionD, inputQuestion.CorrectAnswer)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	if err := eh.db.Create(question).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success created question",
	})
}
