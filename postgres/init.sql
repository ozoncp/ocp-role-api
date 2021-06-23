CREATE DATABASE ocp_role_api;

\connect ocp_role_api

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    service varchar NOT NULL,
    operation varchar NOT NULL
);