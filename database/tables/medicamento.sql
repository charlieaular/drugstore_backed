-- public.medicamento definition

-- Drop table

-- DROP TABLE public.medicamento;

CREATE TABLE public.medicamento (
	id serial4 NOT NULL,
	nombre varchar(50) NULL,
	precio float8 NULL,
	ubicacion varchar(50) NULL,
	CONSTRAINT medicamento_pk PRIMARY KEY (id)
);
CREATE INDEX medicamento_id_idx ON public.medicamento USING btree (id);