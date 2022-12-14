CREATE TABLE categories (
	id SERIAL PRIMARY KEY,
	category VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE goods (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	description TEXT NOT NULL,
	image TEXT,
	category varchar(255) NOT NULL REFERENCES categories (category) ON DELETE CASCADE
);

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name varchar(255),
	username varchar(255) NOT NULL UNIQUE,
	password varchar(80000) NOT NULL,
	role varchar(255) NOT NULL
);
