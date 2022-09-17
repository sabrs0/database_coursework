create or replace function delete_all_foundrisings_of_found() returns trigger as $$
begin
		DELETE FROM foundrising_tab WHERE found_id = OLD.id;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
		DELETE FROM transaction_tab 
		WHERE from_id = OLD.id AND from_essence_type = true;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
		DELETE FROM transaction_tab 
		WHERE to_id = OLD.id AND to_essence_type = false;
		IF NOT FOUND THEN 
			RETURN NULL; 
		END IF;
	RETURN OLD;
end;
$$ LANGUAGE plpgsql;

drop trigger if exists foundrisings_deleting on foundation_tab;

CREATE TRIGGER foundrisings_deleting
AFTER DELETE ON foundation_tab
    FOR EACH ROW EXECUTE PROCEDURE delete_all_foundrisings_of_found();
--\i 'C:/Projects/Go/src/db_course/sql_scripts/trigger.sql'