package model

import (
	"errors"

	"github.com/sanhuanshisanshao/go-zero/core/stores/mon"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
)
