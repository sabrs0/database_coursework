DROP DATABASE IF EXISTS foundations;
CREATE DATABASE foundations;

\c foundations;

CREATE TABLE USER_TAB(
                     id integer NOT NULL PRIMARY KEY,
                    login varchar(100)  NOT NULL,
					password varchar(100)  NOT NULL,
					balance money,
                    charity_Sum money ,
                    closed_Fing_Amount integer
					);

CREATE TABLE foundation_TAB(
                    id integer NOT NULL PRIMARY KEY,
                    name varchar(100)  NOT NULL,
					password varchar(100)  NOT NULL,
                    cur_Foudrising_Amount integer,
                    closed_Foundrising_Amount integer,
                    fund_balance money,
                    income_history money,
                    outcome_history money,
                    volunteer_amount integer);
					login varchar(100)  NOT NULL,
					
					
CREATE TABLE foundrising_TAB(
                    id integer NOT NULL PRIMARY KEY,
                    found_id integer NOT NULL,
                    description text,	
                    required_sum money,
                    current_sum money,
                    philantrops_amount integer,
                    creation_date date,
					closing_date date);
					
CREATE TABLE transaction_TAB(
                    id integer NOT NULL PRIMARY KEY,
					from_essence_type boolean, -- 0 - user 1 - foundation
                    from_id integer NOT NULL,
                    to_essence_type boolean, -- 0 - foundation, 1 - foundrising
                    sum_of_money money,
                    comment text,
                    to_id integer NOT NULL);
					
-- \i 'C:/Projects/Go/src/db_course/sql_scripts/create.sql'