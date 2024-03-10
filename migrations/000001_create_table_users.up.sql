CREATE TABLE  IF NOT EXISTS public.users (
      id bigserial NOT NULL,
      "name" varchar NULL,
      occupation varchar NULL,
      email varchar NOT NULL,
      "password" varchar NULL,
      image varchar NULL,
      image_token varchar NULL,
      "role" varchar NULL,
      created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
      CONSTRAINT users_email_unique UNIQUE (email),
      CONSTRAINT users_pk PRIMARY KEY (id)
);