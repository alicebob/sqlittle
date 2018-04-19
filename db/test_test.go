package db

import (
	"github.com/davecgh/go-spew/spew"
)

func init() {
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.SortKeys = true
}
