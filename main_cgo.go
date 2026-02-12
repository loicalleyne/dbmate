//go:build cgo
// +build cgo

package main

import (
	_ "github.com/loicalleyne/dbmate/v2/pkg/driver/duckdb"
	_ "github.com/loicalleyne/dbmate/v2/pkg/driver/sqlite"
)
