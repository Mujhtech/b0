CREATE TABLE IF NOT EXISTS projects (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

	owner_id uuid NOT NULL REFERENCES users (id),
	name TEXT NOT NULL,
    description TEXT NULL DEFAULT NULL,
    slug TEXT NOT NULL,


	metadata jsonb NOT NULL DEFAULT '{}'::jsonb,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);