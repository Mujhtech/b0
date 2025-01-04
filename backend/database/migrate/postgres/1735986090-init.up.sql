CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE authentication_method AS ENUM ('google', 'github', 'password');

CREATE TYPE token_type AS ENUM ('bearer', 'basic', 'api_key');