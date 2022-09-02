--группировка по кол-ву сборов в каждом фонде
select found_id,count(found_id)
from foundrising_tab
group by found_id
order by found_id
-------------------
update foundation_tab
set curFoudrisingAmount = (		select count(found_id)
								from foundrising_tab
								group by found_id
								order by found_id)
-------
update foundation_tab
set income_history = fund_balance