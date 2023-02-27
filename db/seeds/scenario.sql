-- 初期値の設定
INSERT INTO priorities ("name") VALUES ('high'), ('mid'), ('low');

-- user1が未完了のhigh, mid, low 完了済みのhigh, mid, low 計6個todoを作成
INSERT INTO todos ("id", "title", "description", "priority", "done", "auth0_id") VALUES ('11f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo1', 'description1', 'high', false, 'auth0|01a11f28aacda01526480ddc');
INSERT INTO todos ("id", "title", "description", "priority", "done", "auth0_id") VALUES ('21f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo2', 'description2', 'mid', false, 'auth0|01a11f28aacda01526480ddc');
INSERT INTO todos ("id", "title", "description", "priority", "done", "auth0_id") VALUES ('31f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo3', 'description3', 'low', false, 'auth0|01a11f28aacda01526480ddc');

INSERT INTO todos ("id", "title", "description", "priority", "done", "auth0_id") VALUES ('41f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo1', 'description1', 'high', true, 'auth0|01a11f28aacda01526480ddc');
INSERT INTO todos ("id", "title", "description", "priority", "done", "auth0_id") VALUES ('51f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo2', 'description2', 'mid', true, 'auth0|01a11f28aacda01526480ddc');
INSERT INTO todos ("id", "title", "description", "priority", "done", "auth0_id") VALUES ('61f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo3', 'description3', 'low', true, 'auth0|01a11f28aacda01526480ddc');

-- 加入したuser2が未完了のhigh, mid, low 3つのタスクをもつ

INSERT INTO todos ("id", "title", "description", "priority", "done", "auth0_id") VALUES ('01b607ef-e183-42c7-8420-0540ce982024', 'todo1', 'description1', 'high', false, 'google-oauth2|032852562224003772841');
INSERT INTO todos ("id", "title", "description", "priority", "done", "auth0_id") VALUES ('02b607ef-e183-42c7-8420-0540ce982024', 'todo1', 'description1', 'mid', false, 'google-oauth2|032852562224003772841');
INSERT INTO todos ("id", "title", "description", "priority", "done", "auth0_id") VALUES ('03b607ef-e183-42c7-8420-0540ce982024', 'todo1', 'description1', 'low', false, 'google-oauth2|032852562224003772841');
