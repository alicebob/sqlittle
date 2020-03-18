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
		rows, err := c.Query("cleary a placeholder SQL statement")
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
}
