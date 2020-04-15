package config

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/ini.v1"
)

type Appconfig struct {
	Name     string
	Apppath  string
	Dirs     []string
	Files    []string
	Template []string
}

const basedir = "/home/steven/go/petdog"

func NewAppconfig(name string) *Appconfig {
	return &Appconfig{Name: name}
}

// Parse parse params from file config/app.ini
func (c *Appconfig) Parse() error {
	params := []string{"dirs", "files", "templates"}
	for _, v := range params {
		if err := c.parseparams(v); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}

// parseparams parse struct Appconfig fields values
func (c *Appconfig) parseparams(pname string) error {
	configPath := path.Join(basedir, "./conf/", c.Name+".ini")
	cfg, err := ini.Load(configPath)
	if err != nil {
		return fmt.Errorf("Load ini file error: %w", err)
	}
	val, err := cfg.Section(c.Name).GetKey(pname)
	if err != nil {
		return fmt.Errorf("get ini %s error: %w", pname, err)
	}
	list := val.Strings(",")
	for _, val := range list {
		switch pname {
		case "dirs":
			c.Dirs = append(c.Dirs, val)
		case "files":
			c.Files = append(c.Files, val)
		case "templates":
			c.Template = append(c.Template, val)
		}
	}
	return nil
}

// ParserTemplate parse the template and create file
func (c *Appconfig) ParserTemplate() error {
	if c.Template != nil {
		for _, v := range c.Template {
			temppath := filepath.Join(basedir, "./template", c.Name, v)
			vv := strings.TrimSuffix(v, ".tpl")
			destpath := filepath.Join(c.Apppath, vv)
			if err := c.copyfile(temppath, destpath); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return nil
}

func (c *Appconfig) copyfile(src, dst string) error {
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

	if err := t.Execute(dst, c.Name); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}
	return nil
}
