BEGIN;

CREATE TABLE "todos" (
                         "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                         "title" varchar(80) NOT NULL,
                         "description" varchar,
                         "done" bool NOT NULL DEFAULT false,
                         "priority" varchar(20) NOT NULL DEFAULT 'mid',
                         "due_date" date,
    -- auth0のid
                         "auth0_id" text NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT 'now()',
                         "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "priorities" (
                              "name" varchar(20) PRIMARY KEY,
                              "created_at" timestamptz NOT NULL DEFAULT 'now()',
                              "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "todos" ADD FOREIGN KEY ("priority") REFERENCES "priorities" ("name") ON DELETE SET NULL ON UPDATE CASCADE;

COMMIT;
