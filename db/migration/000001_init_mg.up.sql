
-- Table: public.books


CREATE TABLE IF NOT EXISTS public.books
(
    book_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    book_name text COLLATE pg_catalog."default" NOT NULL,
    book_url text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Books_pkey" PRIMARY KEY (book_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.books
    OWNER to postgres;

GRANT ALL ON TABLE public.books TO "goServer";

GRANT ALL ON TABLE public.books TO postgres;



-- Table: public.chapters

CREATE TABLE IF NOT EXISTS public.chapters
(
    chapter_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    chapter_name text COLLATE pg_catalog."default" NOT NULL,
    chapter_content text COLLATE pg_catalog."default" NOT NULL,
    book_key integer NOT NULL,
    CONSTRAINT "Chapters_pkey" PRIMARY KEY (chapter_id),
    CONSTRAINT book_key FOREIGN KEY (book_key)
        REFERENCES public.books (book_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.chapters
    OWNER to postgres;

GRANT ALL ON TABLE public.chapters TO "goServer";

GRANT ALL ON TABLE public.chapters TO postgres;