package sqlittle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColumns(t *testing.T) {
	db, err := Open("testdata/words.sqlite")
	require.NoError(t, err)
	defer db.Close()

	t.Run("fine", func(t *testing.T) {
		cols, err := db.Columns("words")
		require.NoError(t, err)
		require.Equal(t, []string{"word", "length"}, cols)
	})

	t.Run("no such table", func(t *testing.T) {
		cols, err := db.Columns("notwords")
		require.EqualError(t, err, `no such table: "notwords"`)
		require.Nil(t, cols)
	})
}
