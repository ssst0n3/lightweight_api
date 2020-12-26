package test_data

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_db"
)

func CreateTables(c lightweight_db.Connector) {
	awesome_error.CheckFatal(CreateTableChallenge(c))
	return
}

func CreateTableChallenge(c lightweight_db.Connector) (err error) {
	query := `CREATE TABLE IF NOT EXISTS challenge (
	id integer primary key autoincrement, 
	name text, 
	score integer, 
	solved integer
);`
	statement, err := c.DB.Prepare(query)
	if err != nil {
		awesome_error.CheckFatal(err)
		return
	}
	_, err = statement.Exec()
	awesome_error.CheckErr(err)
	return
}
