CREATE Table Books (
	id BIGSERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL, 
	author VARCHAR(255) NOT NULL,
	publish_date TIMESTAMP not null default now(),
	rating INT NOT NULL
)