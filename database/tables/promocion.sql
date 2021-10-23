-- public.promocion definition

-- Drop table

-- DROP TABLE public.promocion;

CREATE TABLE public.promocion (
	id serial4 NOT NULL,
	descripcion varchar(100) NULL,
	porcentaje float8 NULL,
	fecha_inicio date NULL,
	fecha_fin date NULL,
	CONSTRAINT promocion_pk PRIMARY KEY (id),
	CONSTRAINT promocion_un UNIQUE (fecha_inicio)
);
CREATE INDEX promocion_id_idx ON public.promocion USING btree (id);