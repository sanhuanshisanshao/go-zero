package template

// Error defines an error template
const Error = `package {{.pkg}}

import "github.com/sanhuanshisanshao/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
`
