package user

import (
	"time"

	"github.com/bashmohandes/go-askme/answer"
	"github.com/bashmohandes/go-askme/model"
	"github.com/bashmohandes/go-askme/question"
)

type userUsecase struct {
	questionRepo question.Repository
	answerRepo   answer.Repository
}

// Feed type
type Feed struct {
	Items []*FeedItem
}

// FeedItem type
type FeedItem struct {
	Question   string
	Answer     string
	AnsweredAt time.Time
	Likes      uint
	User       string
}

// AsksUsecase type
type AsksUsecase interface {
	Ask(from *models.User, to *models.User, question string) *models.Question
	Like(user *models.User, answer *models.Answer) uint
	Unlike(user *models.User, answer *models.Answer) uint
	LoadUserFeed(user *models.User) *Feed
}

// AnswersUsecase for the registered user
type AnswersUsecase interface {
	FetchUnansweredQuestions(userID models.UniqueID) []*models.Question
	Answer(user *models.User, question *models.Question, answer string) *models.Answer
}

// NewAsksUsecase creates a new service
func NewAsksUsecase(qRepo question.Repository, aRepo answer.Repository) AsksUsecase {
	return &userUsecase{
		questionRepo: qRepo,
		answerRepo:   aRepo,
	}
}

// NewAnswersUsecase creates a new service
func NewAnswersUsecase(qRepo question.Repository, aRepo answer.Repository) AnswersUsecase {
	return &userUsecase{
		questionRepo: qRepo,
		answerRepo:   aRepo,
	}
}

// LoadQuestions load questions model
func (svc *userUsecase) FetchUnansweredQuestions(userID models.UniqueID) []*models.Question {
	return svc.questionRepo.LoadUnansweredQuestions(userID)
}

// Saves question
func (svc *userUsecase) Ask(from *models.User, to *models.User, question string) *models.Question {
	q := from.Ask(to, question)
	svc.questionRepo.Add(q)
	return q
}

// Likes a question
func (svc *userUsecase) Like(user *models.User, answer *models.Answer) uint {
	svc.answerRepo.AddLike(answer, user)
	return svc.answerRepo.GetLikesCount(answer)
}

// Unlikes a question
func (svc *userUsecase) Unlike(user *models.User, answer *models.Answer) uint {
	svc.answerRepo.RemoveLike(answer, user)
	return svc.answerRepo.GetLikesCount(answer)
}

func (svc *userUsecase) Answer(user *models.User, question *models.Question, answer string) *models.Answer {
	a := user.Answer(question, answer)
	svc.answerRepo.Add(a)
	return a
}

func (svc *userUsecase) LoadUserFeed(user *models.User) *Feed {
	return &Feed{
		Items: []*FeedItem{
			&FeedItem{
				Question:   "What is your name?",
				Answer:     "Mohamed Elsherif daaah!",
				AnsweredAt: time.Now(),
				Likes:      10,
				User:       "Anonymous",
			},
			&FeedItem{
				Question:   "What is your age?",
				Answer:     "I would never tell, but it is 35",
				AnsweredAt: time.Now(),
				Likes:      1,
				User:       "Sayed",
			},
		},
	}
}
