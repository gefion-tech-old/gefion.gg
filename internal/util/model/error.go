package model

type Error struct {
	_type   string
	message string
}

type GGError struct {
	_error Error
}
