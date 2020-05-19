package api

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type Maker struct {
	Temppath string
	Destdir  string
	Filename string
}

const basedir = "/home/steven/go/petdog"

// NewMaker generate new struct Maker
func NewMaker(tempPath, filename string) (Maker, error) {
	var maker Maker
	currPath, err := os.Getwd()
	if err != nil {
		return Maker{}, nil
	}
	if strings.Contains(tempPath, ":") {
		tempname := strings.Split(tempPath, ":")[0]
		maker.Temppath = filepath.Join(basedir, tempname+".go.tpl")
		dstdir := strings.Split(tempPath, ":")[1]
		maker.Destdir = filepath.Join(currPath, dstdir)
	} else {
		maker.Temppath = filepath.Join(basedir, "./conf/template/api.go.tpl")
		maker.Destdir = filepath.Join(currPath, tempPath)
	}
	maker.Filename = filepath.Join(maker.Destdir, filename+".go")
	return maker, nil
}

func (c *Maker) MakeFile() error {
	if c.Temppath == "" {
		return errors.New("No specified template path")
	}
	if c.Destdir == "" {
		return errors.New("no specified file path")
	}

	if err := c.handlerTemplate(c.Temppath, c.Filename); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (c *Maker) handlerTemplate(src, dst string) error {
	tempname := filepath.Base(src)
	t, err := template.New(tempname).ParseFiles(src)
	if err != nil {
		return fmt.Errorf("parse template %s error: %w", src, err)
	}
	f, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("handlerTemplate openfile error: %w", err)
	}
	m := make(map[string]string)
	packagename := filepath.Base(c.Filename)
	m["packagename"] = packagename
	if err := t.Execute(f, m); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}
	return nil
}
