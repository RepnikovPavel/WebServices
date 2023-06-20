package NodeStorage

import (
	"database/sql"
	"time"
)

const (
	SQL_Schema_AccountCold string = `
	CREATE TABLE IF NOT EXISTS cold_account_info(
		id INT NOT NULL,
		time_of_registration TIMESTAMP,
		PRIMATY KEY (id)
	);
	`
	SQL_Insert_AccountCold string = `
	INSERT INTO cold_account_info(
		 time_of_registration
	) VALUES(
		?
	);
	`

	SQL_Schema_AccountHot string = `
	CREATE TABLE IF NOT EXISTS hot_account_info(
		id INT NOT NULL,
		time_of_last_activity TIMESTAMP,
		PRIMATY KEY (id)
	);
	`
	SQL_Schema_StatWarmAccount string = `
	CREATE TABLE IF NOT EXISTS warm_account_stat(
		id INT NOT NULL,
		time_of_last_session TIMESTAMP,
		number_of_fallen_trees INT,
		number_of_purchases INT,
		PRIMATY KEY (id)
	);
	`
	SQL_Schema_HotTmpUserSession string = `
	CREATE TABLE IF NOT EXISTS hot_user_session(
		account_id INT NOT NULL,
		time_of_last_action TIMESTAMP,
		number_of_nails_hammered INT,
		session_duration_second INT,
		number_of_actions INT
	);
	`
)

type SQLiteConf struct {
	db   sql.DB
	time time.Time
}
