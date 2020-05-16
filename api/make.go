package api

import (
	"strings"

	"github.com/pkg/errors"
)

type Maker struct {
	Temppath string
	Destdir  string
	Filename string
}

// NewMaker generate new struct Maker
func NewMaker(tempPath, filename string) (Maker, error) {
	var maker Maker
	if strings.Contains(tempPath, ":") {
		maker.Temppath = strings.Split(tempPath, ":")[0]
		maker.Destdir = strings.Split(tempPath, ":")[1]
	} else {
		maker.Destdir = tempPath
	}
	maker.Filename = filename
	return maker, nil
}

func (c *Maker) ParserTemplate() error {
	if c.Temppath == "" {
		return errors.New("No specified template path")
	}

	return nil
}
