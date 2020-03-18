package driver

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDriver(t *testing.T) {
	c, err := sql.Open("sqlittle", "sqlittle://foobar")
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
		require.Equal(t, []string{"test", "columns"}, cols)

		var res [][]string
		for rows.Next() {
			r := make([]string, 2)
			require.NoError(t, rows.Scan(&r[0], &r[1]))
			res = append(res, r)
		}
		require.NoError(t, rows.Err())
		require.NoError(t, rows.Close())
		require.Equal(t, [][]string{
			{"aap", "noot"},
			{"mies", "wim"},
			{"vuur", "eekhoorn"},
		}, res)
	})
}
