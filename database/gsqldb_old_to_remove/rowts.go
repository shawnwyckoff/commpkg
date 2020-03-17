package gsqldb_old_to_remove

import "time"

func (c *Conn) ReadRowsBefore(table string, tm time.Time) (*Dataset, error) {
	if len(c.connInfo.Database) == 0 {
		return nil, ErrCurrDatabaseIsNull
	}

	return nil, nil
}

func (c *Conn) ReadRowsAfter(table string, tm time.Time) (*Dataset, error) {
	return nil, nil
}

func (c *Conn) ReadRowsBetween(table string, begin, end time.Time) (*Dataset, error) {
	return nil, nil
}

func (c *Conn) RemoveRowsBefore(table string, tm time.Time) error {
	return nil
}

func (c *Conn) RemoveRowsAfter(table string, tm time.Time) error {
	return nil
}

func (c *Conn) RemoveRowsBetween(table string, begin, end time.Time) error {
	return nil
}
