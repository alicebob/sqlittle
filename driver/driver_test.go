package driver

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDriver(t *testing.T) {
	c, err := sql.Open("sqlittle", "../testdata/music.sqlite")
	require.NoError(t, err)
	require.NotNil(t, c)

	t.Run("exec", func(t *testing.T) {
		res, err := c.Exec("foobar")
		require.NoError(t, err)
		aff, err := res.RowsAffected()
		require.NoError(t, err)
		require.EqualValues(t, 0, aff)
	})

	t.Run("query", func(t *testing.T) {
		t.Skip("borked")
		rows, err := c.Query(`SELECT * FROM albums`)
		require.NoError(t, err)
		cols, err := rows.Columns()
		require.NoError(t, err)
		require.Equal(t, []string{"id", "artist", "name"}, cols)

		var res [][]string
		for rows.Next() {
			r := make([]string, 3)
			require.NoError(t, rows.Scan(&r[0], &r[1], &r[2]))
			res = append(res, r)
		}
		require.NoError(t, rows.Err())
		require.NoError(t, rows.Close())
		require.Equal(t, [][]string{
			{"1", "1", "Rubber Soul"},
			{"2", "1", "Abbey Road"},
		}, res)
	})

	t.Run("query, nonrowid table", func(t *testing.T) {
		rows, err := c.Query(`SELECT * FROM tracks`)
		require.NoError(t, err)
		cols, err := rows.Columns()
		require.NoError(t, err)
		require.Equal(t, []string{"id", "album", "name", "length"}, cols)

		var res [][]string
		for rows.Next() {
			r := make([]string, 4)
			require.NoError(t, rows.Scan(&r[0], &r[1], &r[2], &r[3]))
			res = append(res, r)
		}
		require.NoError(t, rows.Err())
		require.NoError(t, rows.Close())
		require.Equal(t, [][]string{
			{"1", "1", "Drive My Car", "145"},
			{"2", "1", "Norwegian Wood", "121"},
			{"3", "1", "You Wont See Me", "198"},
			{"4", "2", "Come Together", "259"},
			{"5", "2", "Something", "182"},
			{"6", "2", "Maxwells Silver Hammer", "207"},
		}, res)
	})

	t.Run("query, non-existing table", func(t *testing.T) {
		rows, err := c.Query(`SELECT * FROM nosuch`)
		require.EqualError(t, err, `no such table: "nosuch"`)
		require.Nil(t, rows)
	})

	t.Run("query, index name used as table", func(t *testing.T) {
		// meta: it would be fun to allow this, though.
		rows, err := c.Query(`SELECT * FROM albums_name`)
		require.EqualError(t, err, `no such table: "albums_name"`)
		require.Nil(t, rows)
	})

	t.Run("query, invalid syntax", func(t *testing.T) {
		rows, err := c.Query(`SELECT * FROM nosuch bar bar`)
		require.EqualError(t, err, `syntax error`)
		require.Nil(t, rows)
	})

	t.Run("query, not a SELECT", func(t *testing.T) {
		rows, err := c.Query(`CREATE TABLE foo (bar INTEGER)`)
		require.EqualError(t, err, `only SELECT is supported (we got a sql.CreateTableStmt)`)
		require.Nil(t, rows)
	})
}
