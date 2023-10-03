package postgres

func (db *DB) GetPermission(pubUserID int, subUserID int) (bool, error) {
	var id int
	if err := db.db.QueryRow("select id from permissions where pub_id = $1 and sub_id = $2", pubUserID, subUserID).Scan(&id); err != nil {
		return false, err
	}
	return true, nil
}

func (db *DB) AddPermission(pubUserID int, subUserID int) (bool, error) {
	if err := db.db.QueryRow("insert into permissions(pub_id, sub_id) values  ($1, $2)", pubUserID, subUserID).Err(); err != nil {
		return false, err
	}
	return true, nil
}
