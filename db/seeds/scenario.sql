-- 初期値の設定
INSERT INTO priorities ("name") VALUES ('high'), ('mid'), ('low');

--  cli用のuser
INSERT INTO users ("id", "auth0_id", "name", "email") VALUES ('1535c2ec-16e4-467d-8120-8ed642bbf7e7', 'iKbdoXOX3b6Bjg5BFnKKzUCgWzq3ic3w@clients','client','client@gmail.com');
-- ブラウザ用のuser
INSERT INTO users ("id", "auth0_id", "name", "email") VALUES ('2535c2ec-16e4-467d-8120-8ed642bbf7e7', 'google-oauth2|032852562224003772841','daisuke','daisukenoguemon@gmail.com');

-- cli用userがownerのproject
INSERT INTO projects("id", "name", "owner_id") VALUES ('d125b07b-9479-4adb-ab5d-1ac5cf8d27f8', 'project1', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
-- ブラウザuserがownerのプロジェクト
INSERT INTO projects("id", "name", "owner_id") VALUES ('d225b07b-9479-4adb-ab5d-1ac5cf8d27f8', 'project2', '2535c2ec-16e4-467d-8120-8ed642bbf7e7');

-- user1が未完了のhigh, mid, low 完了済みのhigh, mid, low 計6個todoを作成
-- 現状memberは存在できない
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('11f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo1', 'description1', 'high', false, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('21f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo2', 'description2', 'mid', false, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('31f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo3', 'description3', 'low', false, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');

INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('41f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo1', 'description1', 'high', true, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('51f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo2', 'description2', 'mid', true, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
INSERT INTO todos ("id", "title", "description", "priority", "done", "project_id", "user_id") VALUES ('61f506ef-db2e-4cc1-a0e7-c7a068cd06bf', 'todo3', 'description3', 'low', true, 'd125b07b-9479-4adb-ab5d-1ac5cf8d27f8', '1535c2ec-16e4-467d-8120-8ed642bbf7e7');
