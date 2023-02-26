CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "user" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "username" VARCHAR(255) UNIQUE,
    "first_name" VARCHAR(255),
    "last_name" VARCHAR(255),
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "password_hash" TEXT NOT NULL, 
    "created_at" INT NOT NULL
);


CREATE TABLE IF NOT EXISTS "url" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id" UUID,
    "created_at" INT NOT NULL,
    "org_path" TEXT NOT NULL,
    "short_hash" TEXT NOT NULL,
    "counter" INT,
);

ALTER TABLE "url" 
ADD CONSTRAINT "fk_url_user_id" 
FOREIGN KEY ("user_id") 
REFERENCES "user"("id") ON DELETE CASCADE;