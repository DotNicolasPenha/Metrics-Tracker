package datawork

import "errors"

type DataWorkHandler struct {
	pathToWork string
}

func NewDataWorkHandler(pathtowork string) (*DataWorkHandler, error) {
	if pathtowork == "" {
		return nil, errors.New("pathToWork is empty")
	}
	return &DataWorkHandler{
		pathToWork: pathtowork,
	}, nil
}
