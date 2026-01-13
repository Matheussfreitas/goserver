package domain

type Question struct {
	ID            string   `json:"id"`
	QuizID        string   `json:"quiz_id"`
	Statement     string   `json:"statement"`
	Answers       []string `json:"answers"`        // Lista de opções (ex: Paris, Londres)
	CorrectAnswer int      `json:"correct_answer"` // Índice da resposta correta
	Explanation   string   `json:"explanation"`
}
