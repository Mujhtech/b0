CREATE TABLE IF NOT EXISTS code_versions (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

	owner_id uuid NOT NULL REFERENCES users (id),
    project_id uuid NOT NULL REFERENCES projects (id),
    endpoint_id uuid NOT NULL REFERENCES endpoints (id),
    version TEXT NOT NULL,
    branch TEXT NOT NULL,
    commit_msg TEXT NOT NULL,
    content jsonb NOT NULL DEFAULT '{}'::jsonb,
	
	metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);