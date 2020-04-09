package gsqldb_old_to_remove

// 如果pgsql有特殊处理，请放在这里
/*
import (
	"github.com/jackc/pgx"
	"strings"
)

func _pgIsDatabaseExists(connectString, database string) (bool, error) {
	cfg, err := pgx.ParseConnectionString(connectString)
	if err != nil {
		return false, err
	}
	cfg.Database = database
	conn, err := pgx.Connect(cfg)

	// Database exists.
	if err == nil {
		conn.C()
		return true, nil
	}
	// Database doesn't exist.
	if strings.Contains(err.Error(), "SQLSTATE 3D000") {
		return false, nil
	}
	// Other errors occurred.
	return false, err
}
*/
