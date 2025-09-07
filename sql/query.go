package sql

import _ "embed"
//go:embed query/insert_word.sql
var InsertWord string
//go:embed query/update_word.sql
var UpdateWord string
