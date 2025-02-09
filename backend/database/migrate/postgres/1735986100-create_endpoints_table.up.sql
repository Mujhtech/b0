CREATE TABLE IF NOT EXISTS endpoints (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),

	owner_id uuid NOT NULL REFERENCES users (id),
    project_id uuid NOT NULL REFERENCES projects (id),
	name TEXT NOT NULL,
    path TEXT NOT NULL,
	method endpoint_method NOT NULL DEFAULT 'GET',
	description TEXT NULL DEFAULT NULL,
	connectors jsonb NOT NULL DEFAULT '{}'::jsonb,
	is_public BOOLEAN NOT NULL DEFAULT FALSE,
	status endpoint_status NOT NULL DEFAULT 'draft',

	workflows jsonb NOT NULL DEFAULT '{}'::jsonb,
	code_generation jsonb NOT NULL DEFAULT '{}'::jsonb,
	metadata jsonb NOT NULL DEFAULT '{}'::jsonb,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);