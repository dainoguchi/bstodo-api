BEGIN;

CREATE TABLE "users" (
                         "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                         "auth0_id" varchar(40) UNIQUE NOT NULL,
                         "name" varchar(20) NOT NULL,
                         "email" varchar(100) UNIQUE NOT NULL,
                         "picture" varchar(255),
                         "created_at" timestamptz NOT NULL DEFAULT 'now()',
                         "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "todos" (
                         "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                         "title" varchar(80) NOT NULL,
                         "description" varchar,
                         "done" bool NOT NULL DEFAULT false,
                         "priority" varchar(20) NOT NULL DEFAULT 'mid',
                         "due_date" timestamptz,
                         "project_id" uuid NOT NULL,
                         "user_id" uuid NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT 'now()',
                         "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "priorities" (
                              "priority" varchar(20) PRIMARY KEY,
                              "created_at" timestamptz NOT NULL DEFAULT 'now()',
                              "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);


CREATE TABLE "projects" (
                            "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                            "owner_id" uuid NOT NULL,
                            "name" varchar(40) UNIQUE NOT NULL,
                            "invitation_token" varchar(40) UNIQUE NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT 'now()',
                            "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "project_users" (
                                 "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                                 "project_id" uuid NOT NULL,
                                 "user_id" uuid NOT NULL,
                                 "role" varchar(20) NOT NULL DEFAULT 'editor',
                                 "created_at" timestamptz NOT NULL DEFAULT 'now()',
                                 "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "user_roles" (
                              "role" varchar(20) PRIMARY KEY,
                              "created_at" timestamptz NOT NULL DEFAULT 'now()',
                              "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE UNIQUE INDEX ON "project_users" ("project_id", "user_id");

ALTER TABLE "project_users" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "project_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "todos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "todos" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "projects" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE "project_users" ADD FOREIGN KEY ("role") REFERENCES "user_roles" ("role") ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE "todos" ADD FOREIGN KEY ("priority") REFERENCES "priorities" ("priority") ON DELETE SET NULL ON UPDATE CASCADE;

COMMIT;