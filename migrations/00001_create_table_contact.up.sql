CREATE TABLE public.contact
(
   id bigserial NOT NULL,
   PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.contact
   OWNER to appuser;