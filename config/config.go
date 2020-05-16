package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"gopkg.in/ini.v1"
)

type Appconfig struct {
	Basedir string
	Name    string
	Apppath string
	Dirs    []string
}

func (c *Appconfig) getBasedir() error {
	//filePath := filePath.Join(c.)
	cfg, err := ini.Load("./conf/app.conf")
	if err != nil {
		return fmt.Errorf("Load ini file error: %w", err)
	}
	val, err := cfg.Section("").GetKey("basedir")
	if err != nil {
		return fmt.Errorf("get ini basedir info error: %w", err)
	}

	c.Basedir = val.Value()
	return nil
}

func NewAppconfig(name string) *Appconfig {
	return &Appconfig{Name: name}
}

// Parse parse params from file config/app.ini
func (c *Appconfig) Parse() error {
	err := c.getBasedir()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	currentPath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	c.Apppath = currentPath

	if err := c.parseparams(); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// parseparams parse struct Appconfig fields values
func (c *Appconfig) parseparams() error {
	configPath := path.Join(c.Basedir, "./conf/", c.Name+".ini")
	cfg, err := ini.Load(configPath)
	if err != nil {
		return fmt.Errorf("Load ini file error: %w", err)
	}
	val, err := cfg.Section(c.Name).GetKey("dirs")
	if err != nil {
		return fmt.Errorf("get ini dirs error: %w", err)
	}
	list := val.Strings(",")
	for _, val := range list {
		c.Dirs = append(c.Dirs, val)
	}
	return nil
}

func (c *Appconfig) GenerateMain() error {
	src := filepath.Join(c.Basedir, "./conf/template/main.go.tpl")

	dst := filepath.Join(c.Apppath, "main.go")
	err := c.copyFile(src, dst)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (c *Appconfig) copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("read template file error: %w", err)
	}

	err = ioutil.WriteFile(dst, input, 0644)
	if err != nil {
		return fmt.Errorf("copy template file error: %w", err)
	}
	return nil
}

func (c *Appconfig) handlerTemplate(src, dst string) error {
	tempname := filepath.Base(src)
	t, err := template.New(tempname).ParseFiles(src)
	if err != nil {
		return fmt.Errorf("parse template %s error: %w", src, err)
	}
	f, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("handlerTemplate openfile error: %w", err)
	}
	if err := t.Execute(f, c.Name); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}
	return nil
}
