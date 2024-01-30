CREATE Table Books (
	id BIGSERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL, 
	author VARCHAR(255) NOT NULL,
	publish_date TIMESTAMP not null default now(),
	rating INT NOT NULL
)

CREATE Table Users (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL, 
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	registered_at TIMESTAMP not null
)