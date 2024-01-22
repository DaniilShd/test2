CREATE TABLE public.persons (
	id serial4 NOT NULL,
	"name" varchar NULL,
	surname varchar NULL,
	patronymic varchar NULL,
	gender varchar NULL,
	nationality varchar NULL,
	age int4 NULL,
	CONSTRAINT persons_pk PRIMARY KEY (id)
);