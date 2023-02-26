CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "username" VARCHAR(255) UNIQUE,
    "first_name" VARCHAR(255),
    "last_name" VARCHAR(255),
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "password_hash" TEXT NOT NULL, 
    "created_at" INT NOT NULL
);


CREATE TABLE IF NOT EXISTS "urls" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id" UUID,
    "org_path" TEXT NOT NULL,
    "short_path" TEXT NOT NULL,
    "counter" INT,
    "created_at" INT NOT NULL,
    "updated_at" INT NOT NULL,
    "qrcode_path" TEXT,
    "status" INT NOT NULL
);

ALTER TABLE "urls" 
ADD CONSTRAINT "fk_url_user_id" 
FOREIGN KEY ("user_id") 
REFERENCES "users"("id") ON DELETE CASCADE;