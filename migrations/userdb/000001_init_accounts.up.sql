CREATE TABLE IF NOT EXISTS public.account
(
    id bigserial NOT NULL,
    username character varying(50) NOT NULL DEFAULT '',
    password character varying(50) NOT NULL DEFAULT '',
    enabled boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id),
    UNIQUE (username)
);