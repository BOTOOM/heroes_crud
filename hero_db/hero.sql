-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.3-beta1
-- PostgreSQL version: 9.5
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.
-- 
-- object: new_database | type: DATABASE --
-- DROP DATABASE IF EXISTS new_database;
CREATE DATABASE new_database;
-- ddl-end --


-- object: public.publisher | type: TABLE --
-- DROP TABLE IF EXISTS public.publisher CASCADE;
CREATE TABLE public.publisher (
	id serial NOT NULL,
	name varchar(50) NOT NULL,
	active boolean,
	create_in timestamp,
	CONSTRAINT publisher_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.publisher OWNER TO postgres;
-- ddl-end --

-- object: public.hero | type: TABLE --
-- DROP TABLE IF EXISTS public.hero CASCADE;
CREATE TABLE public.hero (
	id serial NOT NULL,
	name varchar(100) NOT NULL,
	age smallint,
	money decimal(20,2),
	create_in timestamp,
	publisher_id integer,
	CONSTRAINT hero_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.hero OWNER TO postgres;
-- ddl-end --

-- object: public.power | type: TABLE --
-- DROP TABLE IF EXISTS public.power CASCADE;
CREATE TABLE public.power (
	id serial NOT NULL,
	name varchar(200) NOT NULL,
	create_in timestamp,
	hero_id integer,
	CONSTRAINT power_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.power OWNER TO postgres;
-- ddl-end --

-- object: publisher_fk | type: CONSTRAINT --
-- ALTER TABLE public.hero DROP CONSTRAINT IF EXISTS publisher_fk CASCADE;
ALTER TABLE public.hero ADD CONSTRAINT publisher_fk FOREIGN KEY (publisher_id)
REFERENCES public.publisher (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: hero_fk | type: CONSTRAINT --
-- ALTER TABLE public.power DROP CONSTRAINT IF EXISTS hero_fk CASCADE;
ALTER TABLE public.power ADD CONSTRAINT hero_fk FOREIGN KEY (hero_id)
REFERENCES public.hero (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --


