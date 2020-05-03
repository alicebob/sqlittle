package driver

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDriver(t *testing.T) {
	c, err := sql.Open("sqlittle", "../testdata/music.sqlite")
	require.NoError(t, err)
	require.NotNil(t, c)

	t.Run("exec", func(t *testing.T) {
		_, err := c.Exec("foobar")
		require.EqualError(t, err, "Exec() is not supported")
	})

	t.Run("query", func(t *testing.T) {
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
		require.NoError(t, rows.Close())
		require.NoError(t, rows.Err())
		require.Equal(t, [][]string{
			{"1", "1", "Rubber Soul"},
			{"2", "1", "Abbey Road"},
		}, res)
	})

	t.Run("query, columns", func(t *testing.T) {
		rows, err := c.Query(`SELECT name, id, id FROM albums`)
		require.NoError(t, err)
		cols, err := rows.Columns()
		require.NoError(t, err)
		require.Equal(t, []string{"name", "id", "id"}, cols)

		var res [][]string
		for rows.Next() {
			r := make([]string, 3)
			require.NoError(t, rows.Scan(&r[0], &r[1], &r[2]))
			res = append(res, r)
		}
		require.NoError(t, rows.Close())
		require.NoError(t, rows.Err())
		require.Equal(t, [][]string{
			{"Rubber Soul", "1", "1"},
			{"Abbey Road", "2", "2"},
		}, res)
	})

	t.Run("query, columns with *", func(t *testing.T) {
		rows, err := c.Query(`SELECT * FROM albums`)
		require.NoError(t, err)
		cols, err := rows.Columns()
		require.NoError(t, err)
		require.Equal(t, []string{"id", "artist", "name"}, cols)
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
		require.NoError(t, rows.Close())
		require.NoError(t, rows.Err())
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

	t.Run("query, invalid columns", func(t *testing.T) {
		rows, err := c.Query(`SELECT nope, such FROM albums`)
		// - ideally:
		// require.EqualError(t, err, "invalid col")
		// require.Nil(t, rows)
		// - but for now:
		require.NoError(t, err)
		require.False(t, rows.Next())
		require.NoError(t, rows.Close())
		require.EqualError(t, rows.Err(), `no such column: "nope"`)
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

func TestClose(t *testing.T) {
	c, err := sql.Open("sqlittle", "../testdata/music.sqlite")
	require.NoError(t, err)
	require.NotNil(t, c)

	rows, err := c.Query(`SELECT * FROM albums`)
	require.NoError(t, err)
	cols, err := rows.Columns()
	require.NoError(t, err)
	require.Equal(t, []string{"id", "artist", "name"}, cols)

	// there are two rows. We load only 1 and then Close().
	require.True(t, rows.Next())
	require.NoError(t, rows.Close())
	require.NoError(t, rows.Err())

	require.NoError(t, c.Close())
}

func TestQueryContext(t *testing.T) {
	c, err := sql.Open("sqlittle", "../testdata/music.sqlite")
	require.NoError(t, err)
	require.NotNil(t, c)

	ctx, cancel := context.WithCancel(context.Background())
	rows, err := c.QueryContext(ctx, `SELECT * FROM albums`)
	require.NoError(t, err)
	cols, err := rows.Columns()
	require.NoError(t, err)
	require.Equal(t, []string{"id", "artist", "name"}, cols)

	// there are two rows. We load only 1 and then cancel the context.
	require.True(t, rows.Next())
	cancel()
	require.NoError(t, rows.Err())

	require.NoError(t, c.Close())
}

func TestQueryRow(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c, err := sql.Open("sqlittle", "../testdata/music.sqlite")
		require.NoError(t, err)
		require.NotNil(t, c)

		var name string
		err = c.QueryRow(`SELECT name FROM albums`).Scan(&name)
		require.NoError(t, err)
		require.Equal(t, name, "Rubber Soul")
	})

	t.Run("invalid table", func(t *testing.T) {
		c, err := sql.Open("sqlittle", "../testdata/music.sqlite")
		require.NoError(t, err)
		require.NotNil(t, c)

		var name string
		err = c.QueryRow(`SELECT name FROM nosuchtable`).Scan(&name)
		require.EqualError(t, err, `no such table: "nosuchtable"`)
		require.Equal(t, name, "")
	})

	t.Run("invalid column", func(t *testing.T) {
		c, err := sql.Open("sqlittle", "../testdata/music.sqlite")
		require.NoError(t, err)
		require.NotNil(t, c)

		var name string
		err = c.QueryRow(`SELECT nosuch FROM albums`).Scan(&name)
		require.EqualError(t, err, `no such column: "nosuch"`)
		require.Equal(t, name, "")
	})
}
