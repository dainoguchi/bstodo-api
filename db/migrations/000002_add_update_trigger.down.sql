BEGIN;

drop trigger update_trigger on todos;
drop trigger update_trigger on priorities;
drop trigger update_trigger on projects;
drop trigger update_trigger on users;
-- drop trigger update_trigger on project_users;
-- drop trigger update_trigger on roles;

drop routine if exists set_updated_at();

COMMIT;
