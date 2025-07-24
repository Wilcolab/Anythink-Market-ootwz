-- Create questions table
CREATE TABLE IF NOT EXISTS questions (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    options JSONB NOT NULL,
    answer INTEGER NOT NULL,
    category VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create index on category for faster filtering
CREATE INDEX IF NOT EXISTS idx_questions_category ON questions(category);

-- Insert sample data
INSERT INTO questions (text, options, answer, category) VALUES
('What is the capital of France?', '["London", "Berlin", "Paris", "Madrid"]', 2, 'Geography'),
('Which programming language is known for its simplicity and efficiency?', '["Java", "Go", "C++", "Python"]', 1, 'Programming'),
('What is 2 + 2?', '["3", "4", "5", "6"]', 1, 'Math'),
('Who wrote ''Romeo and Juliet''?', '["Charles Dickens", "William Shakespeare", "Jane Austen", "Mark Twain"]', 1, 'Literature'),
('What is the largest planet in our solar system?', '["Earth", "Jupiter", "Saturn", "Mars"]', 1, 'Science'),
('Which year did World War II end?', '["1944", "1945", "1946", "1947"]', 1, 'History')
ON CONFLICT DO NOTHING;
