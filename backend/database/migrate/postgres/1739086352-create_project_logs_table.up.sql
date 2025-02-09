CREATE TABLE IF NOT EXISTS project_logs (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

	owner_id uuid NOT NULL REFERENCES users (id),
    project_id uuid NOT NULL REFERENCES projects (id),
    endpoint_id uuid NULL DEFAULT NULL REFERENCES endpoints (id),
    log_type TEXT NOT NULL,
    log_data jsonb NOT NULL DEFAULT '{}'::jsonb,
	
	metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);