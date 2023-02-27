BEGIN;

drop trigger update_trigger on todos;
drop trigger update_trigger on priorities;

drop routine if exists set_updated_at();

COMMIT;
