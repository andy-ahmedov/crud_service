CREATE Table Books (
	id BIGSERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL, 
	author VARCHAR(255) NOT NULL,
	publish_date TIMESTAMP not null default now(),
	rating INT NOT NULL
);

CREATE Table Users (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL, 
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	registered_at TIMESTAMP not null
);

CREATE Table refresh_tokens (
	id serial NOT NULL UNIQUE,
	user_id int REFERENCES Users (id) on delete CASCADE NOT NULL,
	token VARCHAR(255) NOT NULL UNIQUE,
	expires_at TIMESTAMP NOT NULL
);