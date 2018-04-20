## SQLite file reader

master plan:

err := db.Select("table", cb, "col1", "col2")
err := db.IndexedSelect("table", "index", cb, "col1", "col2")
row, err := db.SelectRowid("table", rowid, "col1", "col2")
#err := db.PKSelect("table", []{...}, cb, "col1", "col2")
#row, err := db.SelectPK("table", []{...}, cb, "col1", "col2")
#err := db.IndexedSelectEq("table", "index", []{...}, cb, "col1", "col2")
#err := db.IndexedSelectFrom("table", "index", []{...}, cb, "col1", "col2")
##err := db.IndexedSelectTo("table", "index", []{...}, cb, "col1", "col2")
##err := db.IndexedSelectFromTo("table", "index", []{...}, []{...}, cb, "col1", "col2")
