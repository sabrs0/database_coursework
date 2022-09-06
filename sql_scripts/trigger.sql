create or replace function delete_all_foundrisings_of_found() returns trigger as $$
begin
	if (TG_OP = 'DELETE') THEN
		DELETE FROM foundrising_tab WHERE found_id = OLD.id;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
	end if;
	RETURN OLD;
end;
$$ LANGUAGE plpgsql;

CREATE TRIGGER foundrisings_deleting
AFTER DELETE ON foundation_tab
    FOR EACH ROW EXECUTE PROCEDURE delete_all_foundrisings_of_found();
--\i 'C:/Projects/Go/src/db_course/sql_scripts/trigger.sql'