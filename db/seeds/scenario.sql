-- user1がownerのプロジェクトにuser2が参加しtodoを共有する

-- 初期値の設定
INSERT INTO priorities ("name") VALUES ('high'), ('mid'), ('low');
INSERT INTO roles ("role") VALUES ('editor'), ('admin');

-- user1, user2, user3を作成
INSERT INTO users ("id", "auth0_id", "name", "email") VALUES ('1535c2ec-16e4-467d-8120-8ed642bbf7e7', 'user1','user1','user1@gmail.com');
INSERT INTO users ("id", "auth0_id", "name", "email") VALUES ('2535c2ec-16e4-467d-8120-8ed642bbf7e7', 'user2','user2','user2@gmail.com');
INSERT INTO users ("id", "auth0_id", "name", "email") VALUES ('3535c2ec-16e4-467d-8120-8ed642bbf7e7', 'user3','user3','user3@gmail.com');

-- user1がproject1を作成 = project1のオーナーがuser1
INSERT INTO projects("id", "name", "invitation_token", "owner_id") VALUES ('d125b07b-9479-4adb-ab5d-1ac5cf8d27f8', 'project1', 'token', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');

-- user1が未完了のhigh, mid, low 完了済みのhigh, mid, low 計6個todoを作成
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('11f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo1', 'description1', 'high', false, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('21f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo2', 'description2', 'mid', false, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('31f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo3', 'description3', 'low', false, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');

INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('41f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo1', 'description1', 'high', true, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('51f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo2', 'description2', 'mid', true, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('61f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo3', 'description3', 'low', true, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');

-- user2がメンバーとして加入
INSERT INTO project_users ("id", "role", "project_id", "user_id") VALUES ('1509c316-343b-48f5-81d8-393032e4be0f', 'editor', 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '2535c2ec-16e4-467d-8120-8ed642bbf7e7');

-- 加入したuser2が未完了のhigh, mid, low 3つのタスクをもつ
-- id = 01b607ef-e183-42c7-8420-0540ce982024
-- project_id = d125b07b-9479-4adb-ab5d-1ac5cf8d27f8
-- user2のid = 2535c2ec-16e4-467d-8120-8ed642bbf7e7
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('01b607ef-e183-42c7-8420-0540ce982024', 'todo1', 'description1', 'high', false, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '2535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('02b607ef-e183-42c7-8420-0540ce982024', 'todo1', 'description1', 'mid', false, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '2535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('03b607ef-e183-42c7-8420-0540ce982024', 'todo1', 'description1', 'low', false, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '2535c2ec-16e4-467d-8120-8ed642bbf7e7');
