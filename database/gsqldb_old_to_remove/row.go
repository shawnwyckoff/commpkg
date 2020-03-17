package gsqldb_old_to_remove

func (c *Conn) UpsertRow(record interface{}) error {
	if len(c.connInfo.Database) == 0 {
		return ErrCurrDatabaseIsNull
	}

	return c.gormConn.Save(record).Error
}
