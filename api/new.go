package api

import (
	"fmt"
	"gotools/conf"
	"gotools/utils"
	"os"
	path "path/filepath"
	"strings"
)

func NewProject(appname string) error {
	currpath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	appPath := path.Join(currpath, appname)
	var modelname string
	if strings.Contains(appname, ":") {
		modelname = strings.Split(appname, ":")[0]
	} else {
		modelname = "base"
	}

	if utils.IsExist(appname) {
		return fmt.Errorf("Error: the same file or directory name has been exists, -%s", appname)
	}
	if err := os.MkdirAll(appPath, 0755); err != nil {
		return err
	}
	fmt.Printf("mkdir directory %s\n", appPath)
	if err := generateDirAndFile(appPath, modelname); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := os.Chdir(appPath); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := utils.ExecShell("/usr/bin/go", "mod", "init", appname); err != nil {
		return fmt.Errorf("go mod init error: %w", err)
	}

	return nil
}

func generateDirAndFile(appdir, modelname string) error {
	appconfig := conf.NewAppconfig(modelname)
	if err := appconfig.Parse(); err != nil {
		return fmt.Errorf("%w", err)
	}

	for _, dir := range appconfig.Dirs {
		if err := os.Mkdir(path.Join(appdir, dir), 0755); err != nil {
			return err
		}
		fmt.Printf("mkdir directory %s/%s\n", appdir, dir)
	}

	for _, filename := range appconfig.Files {
		filepath := path.Join(appdir, filename)
		f, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("create file error: %w", err)
		}
		f.Close()
		fmt.Printf("created file %s\n", filepath)
	}
	if err := appconfig.ParserTemplate(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
