package dberrs

import "errors"

var ErrorNotUniqueTitle = errors.New("title should be unique")
var ErrorUnknownTimeFilter = errors.New("unknown time filter")
