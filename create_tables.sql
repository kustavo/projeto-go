CREATE TABLE IF NOT EXISTS public.users (
	id int NOT NULL GENERATED ALWAYS AS IDENTITY,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	username varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	"name" varchar(255) NOT NULL,
	email varchar(255) NOT NULL
);
