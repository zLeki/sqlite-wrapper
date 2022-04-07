package newsfeed

import "database/sql"
type Item struct {
	ID int
	Content string
}
type Feed struct {
	DB *sql.DB
	Table string
}
func (feed *Feed) Delete(id int) error {
	_, err := feed.DB.Exec("DELETE FROM "+feed.Table+" WHERE id = ?", id)

	return err
}
func (feed *Feed) Query() []Item {
	items := []Item{}
	rows, _ := feed.DB.Query("SELECT * FROM "+feed.Table)
	var id int
	var content string
	for rows.Next() {

		rows.Scan(&id, &content)
		items = append(items, Item{ID: id, Content: content})

	}
	return items
}
func (feed *Feed) Add(item Item) {
	stmt, err := feed.DB.Prepare("INSERT INTO "+feed.Table+" (content) VALUES (?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(item.Content)
	if err != nil {
		panic(err)
	}
}
func NewTable(db *sql.DB, table string) *Feed {
	stmt ,_ := db.Prepare("CREATE TABLE IF NOT EXISTS "+table+" (id INTEGER PRIMARY KEY, content TEXT)")
	_, err := stmt.Exec()
	if err != nil {
		panic(err)
	}
	return &Feed{DB: db, Table: table}
}

