package quizzes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goserver/internal/config"
	"goserver/internal/repository"
)

type CreateQuizService struct {
	repo *repository.QuizRepository
	db   *sql.DB
}

func NewCreateQuizService(repo *repository.QuizRepository, db *sql.DB) *CreateQuizService {
	return &CreateQuizService{
		repo: repo,
		db:   db,
	}
}

type QuizExpected struct {
	QuizTitle string `json:"quiz_title"`
	Questions []struct {
		Statement    string `json:"statement"`
		Alternatives []struct {
			Text string `json:"text"`
		} `json:"alternatives"`
		CorrectIndex int    `json:"correct_index"`
		Explanation  string `json:"explanation"`
	} `json:"questions"`
}

func (s *CreateQuizService) CreateQuiz(tema string, numQuestoes int, dificuldade string) (string, error) {
	prompt := BuildPrompt(tema, numQuestoes, dificuldade)

	quizBuild, err := config.Gemini(prompt)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var quizExpected QuizExpected

	if err := json.Unmarshal([]byte(quizBuild), &quizExpected); err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(quizExpected)

	return quizExpected.QuizTitle, nil
}

func BuildPrompt(tema string, numQuestoes int, dificuldade string) string {
	return fmt.Sprintf(`Atue como um professor especialista no tema %s. 
    Gere um quiz com %d questões de nível %s.
    Retorne apenas JSON no formato:
    {
      "quiz_title": string,
      "questions": [
        {
          "statement": string,
          "alternatives": [{"text": string}],
          "correct_index": int,
          "explanation": string
        }
      ]
    }
    Não inclua explicações fora do JSON.`, tema, numQuestoes, dificuldade)
}
