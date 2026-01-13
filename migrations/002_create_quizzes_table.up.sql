CREATE TYPE difficulty_level AS ENUM ('easy', 'medium', 'hard');

CREATE TABLE quizzes (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  
	title VARCHAR(255) NOT NULL,
	content TEXT NOT NULL,
	difficulty difficulty_level NOT NULL,
	number_questions INTEGER NOT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);