ALTER TABLE IF EXISTS public.contact
    ADD COLUMN mother_name character varying(100) NOT NULL DEFAULT '';

ALTER TABLE IF EXISTS public.contact
    ADD COLUMN father_name character varying NOT NULL DEFAULT '';