CREATE TABLE IF NOT EXISTS ai_usages (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

	owner_id uuid NOT NULL REFERENCES users (id),
    project_id uuid NOT NULL REFERENCES projects (id),
    endpoint_id uuid NULL DEFAULT NULL REFERENCES endpoints (id),
	input_tokens TEXT NOT NULL,
    output_tokens TEXT NOT NULL,
	model TEXT NULL DEFAULT NULL,
    usage_type TEXT NULL DEFAULT NULL,
	metadata jsonb NOT NULL DEFAULT '{}'::jsonb,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);