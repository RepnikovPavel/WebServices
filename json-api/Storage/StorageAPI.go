package Storage

import (
	"database/sql"
	DS "json-api/DataStructures"
	conf "json-api/GoConf"
	cmd_ui "json-api/cmd_ui"

	_ "github.com/lib/pq"
)

type PostgreError struct {
	err string
}

func (perr PostgreError) Error() string {
	return perr.err
}

// CRUD API
type AccountAPI interface {
	CreateAccount(DS.AccountUID) error
	DeleteAccount(DS.AccountUID) error
	UpdateAccount(DS.AccountUID) error
	GetAccount(DS.AccountUID) (DS.AllAccountInfo, error)
	// TODO: why it is not compiled
	// GetAccount(DS.AccountUID) (DS.AllAccountInfo, error)
}

type PostgreLayer struct {
	db *sql.DB // pool of connections. green thread safe.
}

// CRUD API implementation

func (pl *PostgreLayer) CreateAccount(auid DS.AccountUID) error {
	return nil
}

func (pl *PostgreLayer) DeleteAccount(auid DS.AccountUID) error {
	return nil
}

func (pl *PostgreLayer) UpdateAccount(auid DS.AccountUID) error {
	return nil
}

func (pl *PostgreLayer) GetAccount(auid DS.AccountUID) (DS.AllAccountInfo, error) {
	return DS.AllAccountInfo{}, nil
}

func NewPostgreLayer() (*PostgreLayer, error) {
	db, db_open_err := sql.Open("postgres", conf.PostgreConnString)
	if db_open_err != nil {
		cmd_ui.LOGERRSEQ(db_open_err, PostgreError{err: "sql.Open() fail"})
		return nil, db_open_err
	}
	// is connection alive and retry conn
	ping_err := db.Ping()
	if ping_err != nil {
		cmd_ui.LOGERRSEQ(PostgreError{err: "app side: db.Ping() fail"}, ping_err)
		return nil, ping_err
	}
	pl := &PostgreLayer{
		db: db,
	}
	return pl, nil
}

func (pl *PostgreLayer) CreateORValidateSchema() {

}

func SyncRunSQLQuery() {

}
