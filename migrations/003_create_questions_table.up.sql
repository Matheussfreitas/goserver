CREATE TABLE questions (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	quiz_id UUID NOT NULL REFERENCES quizzes(id) ON DELETE CASCADE,

	statement TEXT NOT NULL,
	answers JSONB NOT NULL,
	correct_answer INTEGER NOT NULL,
	explanation TEXT
);