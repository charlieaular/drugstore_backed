-- public.factura definition

-- Drop table

-- DROP TABLE public.factura;

CREATE TABLE public.factura (
	id serial4 NOT NULL,
	fecha_crear date NULL,
	pago_total float8 NULL,
	promocion_id int4 NULL,
	CONSTRAINT factura_pk PRIMARY KEY (id)
);
CREATE INDEX factura_id_idx ON public.factura USING btree (id);


-- public.factura foreign keys

ALTER TABLE public.factura ADD CONSTRAINT factura_fk FOREIGN KEY (promocion_id) REFERENCES public.promocion(id) ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE public.factura ALTER COLUMN pago_total TYPE decimal(6, 2) USING pago_total::decimal;
