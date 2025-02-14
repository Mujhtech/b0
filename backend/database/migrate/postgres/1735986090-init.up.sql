-- Create DB if not exists
-- CREATE DATABASE IF NOT EXISTS "b0";

-- Assign postgres user to database
-- GRANT ALL PRIVILEGES ON DATABASE "b0" TO postgres;

-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create table enums
CREATE TYPE authentication_method AS ENUM ('google', 'github', 'password');

CREATE TYPE token_type AS ENUM ('bearer', 'basic', 'api_key');

CREATE TYPE endpoint_method AS ENUM ('GET', 'POST', 'PUT', 'PATCH', 'DELETE');

CREATE TYPE endpoint_status AS ENUM ('active', 'inactive', 'draft');