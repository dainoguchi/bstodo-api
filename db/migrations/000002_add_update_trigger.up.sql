CREATE FUNCTION set_updated_at() RETURNS trigger AS $$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_trigger
    BEFORE UPDATE
    ON todos
    FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

CREATE TRIGGER update_trigger
    AFTER UPDATE
    ON priorities
    FOR EACH ROW EXECUTE PROCEDURE set_updated_at();
