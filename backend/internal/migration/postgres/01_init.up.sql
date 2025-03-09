CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" text NOT NULL,
  "avatar_url" text NOT NULL,
  "created_at" timestamp DEFAULT NOW()
);

-- ON DELETE SET "unknown" user
INSERT INTO "users" (
  "id",
  "username",
  "email",
  "password",
  "avatar_url"
) VALUES (
  'e036de42-0284-450d-87d1-cb6e306992f2',
  'unknown',
  'unknown@unknown.com',
  'fe945d98-69eb-4d05-8e07-17ddbea049bf', -- random password that cannot be opened
  'fe945d98-69eb-4d05-8e07-17ddbea049bf' -- TODO: set here default avatar URL from tigres
);

CREATE TABLE IF NOT EXISTS "posts" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" UUID NOT NULL DEFAULT 'e036de42-0284-450d-87d1-cb6e306992f2', -- fk
  "title" VARCHAR NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT now(),
  "rating" INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS "comments" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "post_id" UUID NOT NULL DEFAULT 'e036de42-0284-450d-87d1-cb6e306992f2', -- fk
  "user_id" UUID NOT NULL DEFAULT 'e036de42-0284-450d-87d1-cb6e306992f2', -- fk
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT now()
);

ALTER TABLE "posts" 
ADD CONSTRAINT "posts_users_fk"
FOREIGN KEY ("user_id")
REFERENCES "users"("id")
ON DELETE SET DEFAULT;

ALTER TABLE "comments" 
ADD CONSTRAINT "comments_users_fk"
FOREIGN KEY ("user_id")
REFERENCES "users"("id")
ON DELETE SET DEFAULT;

ALTER TABLE "comments"
ADD CONSTRAINT "comments_posts_fk"
FOREIGN KEY ("post_id")
REFERENCES "posts"("id")
ON DELETE SET DEFAULT;