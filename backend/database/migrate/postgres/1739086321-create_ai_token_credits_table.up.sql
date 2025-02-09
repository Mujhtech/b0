CREATE TABLE IF NOT EXISTS ai_token_credits (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

	owner_id uuid NOT NULL REFERENCES users (id),
	model TEXT NULL DEFAULT NULL,
    credits INT NOT NULL DEFAULT 0,
    total_credits INT NOT NULL DEFAULT 0,
    used_credits INT NOT NULL DEFAULT 0,
	metadata jsonb NOT NULL DEFAULT '{}'::jsonb,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);