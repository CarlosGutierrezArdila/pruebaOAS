-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.1
-- PostgreSQL version: 9.6
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: "pruebaOas" | type: DATABASE --
-- -- DROP DATABASE IF EXISTS "pruebaOas";
-- CREATE DATABASE "pruebaOas"
-- 	OWNER = postgres;
-- -- ddl-end --
-- COMMENT ON DATABASE "pruebaOas" IS 'bd prueba OAS';
-- -- ddl-end --
-- 

-- object: public.actividad | type: TABLE --
-- DROP TABLE IF EXISTS public.actividad CASCADE;
CREATE TABLE public.actividad(
	id_actividad integer NOT NULL,
	f_creaccion date NOT NULL,
	f_limite date NOT NULL,
	descripcion varchar(500) NOT NULL,
	estado varchar(10) NOT NULL,
	id_responsable integer NOT NULL,
	CONSTRAINT pk_actividad PRIMARY KEY (id_actividad)

);
-- ddl-end --
COMMENT ON COLUMN public.actividad.f_limite IS 'fecha limite de la actividad';
-- ddl-end --
ALTER TABLE public.actividad OWNER TO postgres;
-- ddl-end --

-- object: public.responsable | type: TABLE --
-- DROP TABLE IF EXISTS public.responsable CASCADE;
CREATE TABLE public.responsable(
	id_responsable integer NOT NULL,
	nombres varchar(50) NOT NULL,
	apellidos varchar(50) NOT NULL,
	email varchar(50) NOT NULL,
	telefono varchar(50) NOT NULL,
	CONSTRAINT pk_responsable PRIMARY KEY (id_responsable)

);
-- ddl-end --
ALTER TABLE public.responsable OWNER TO postgres;
-- ddl-end --

-- object: responsable | type: CONSTRAINT --
-- ALTER TABLE public.actividad DROP CONSTRAINT IF EXISTS responsable CASCADE;
ALTER TABLE public.actividad ADD CONSTRAINT responsable FOREIGN KEY (id_responsable)
REFERENCES public.responsable (id_responsable) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --


