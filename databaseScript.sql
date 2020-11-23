create extension if not exists "uuid-ossp";

CREATE TABLE public."COLORS" (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"NAME" varchar NOT NULL,
	CONSTRAINT colores_pk PRIMARY KEY (id)
);

CREATE TABLE public."TYPES" (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"NAME" varchar NOT NULL,
	CONSTRAINT types_pk PRIMARY KEY (id)
);

CREATE TABLE public."PRODUCTS" (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"NAME" varchar NOT NULL,
	color uuid NOT NULL,
	"TYPE" uuid NOT NULL,
	CONSTRAINT product_pk PRIMARY KEY (id)
);

ALTER TABLE public."PRODUCTS" ADD CONSTRAINT product_fk FOREIGN KEY (color) REFERENCES "COLORS"(id);
ALTER TABLE public."PRODUCTS" ADD CONSTRAINT product_fk_1 FOREIGN KEY ("TYPE") REFERENCES "TYPES"(id);