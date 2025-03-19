-- Table: public.contact

-- DROP TABLE IF EXISTS public.contact;

CREATE TABLE IF NOT EXISTS public.contact
(
    id bigint NOT NULL DEFAULT nextval('contact_id_seq'::regclass),
    CONSTRAINT contact_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.contact
    OWNER to postgres;