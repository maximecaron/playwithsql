package dbhandler

import "github.com/jmoiron/sqlx"

// ConfType contains the creation of the handler
type ConfType interface {
	NewDBHandler() (*sqlx.DB, error)
}

// Get returns the DB connection
func Get(dbType string, host string, dbName string) (db *sqlx.DB, err error) {

	var conf ConfType

	switch dbType {
	case "mysql":
		conf = &MySQLDB{
			Protocol: "tcp",
			Host:     host,
			Port:     "3306",
			User:     "root",
			Password: "test",
			Dbname:   dbName,
			SSL: SSL{
				CertPath:   "",
				KeyPath:    "",
				CAPath:     "",
				ServerName: "",
			},
		}
	case "sqlite":
		conf = &SQLiteDB{}
	case "postgres":
		conf = &PostgresDB{
			Host:     host,
			Port:     "5432",
			User:     "root",
			Password: "test",
			Dbname:   dbName,
			SSL: SSL{
				CertPath:   "",
				KeyPath:    "",
				CAPath:     "",
				ServerName: "",
			},
		}
	case "cockroachdb":
		conf = &CockroachDB{
			Host:   host,
			Port:   "26257",
			User:   "root",
			Dbname: dbName,
			SSL: SSL{
				CertPath:   "",
				KeyPath:    "",
				CAPath:     "",
				ServerName: "",
			},
		}
	}

	return conf.NewDBHandler()
}
