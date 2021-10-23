-- public.factura_medicamento definition

-- Drop table

-- DROP TABLE public.factura_medicamento;

CREATE TABLE public.factura_medicamento (
	id serial4 NOT NULL,
	factura_id int4 NOT NULL,
	medicamento_id int4 NOT NULL
);


-- public.factura_medicamento foreign keys

ALTER TABLE public.factura_medicamento ADD CONSTRAINT factura_medicamento_fk FOREIGN KEY (factura_id) REFERENCES public.factura(id) ON DELETE RESTRICT ON UPDATE CASCADE;
ALTER TABLE public.factura_medicamento ADD CONSTRAINT factura_medicamento_fk_1 FOREIGN KEY (medicamento_id) REFERENCES public.medicamento(id) ON DELETE RESTRICT ON UPDATE CASCADE;