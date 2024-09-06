package compat

import "github.com/afxcode/go-sqlite"

func init() {
	sqlite.RegisterAsSQLITE3()
}
