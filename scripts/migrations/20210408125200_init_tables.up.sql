CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE sex_type AS ENUM (
    'male',
    'female',
    'other'
    );

CREATE TABLE IF NOT EXISTS users
(
    uuid        uuid DEFAULT uuid_generate_v4(),
    first_name  VARCHAR(30)  NOT NULL,
    last_name   VARCHAR(30)  NOT NULL,
    middle_name VARCHAR(30),
    full_name   VARCHAR(255) NOT NULL,
    sex         sex_type     NOT NULL,
    birth_date  timestamp
);
