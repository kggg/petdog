package api

import (
	"fmt"
	"go/build"
	"os"
	path "path/filepath"
	"petdog/config"
	"petdog/utils"
	"strings"
)

func NewProject(appname string) error {
	currpath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	var modelname string
	if strings.Contains(appname, ":") {
		modelname = strings.Split(appname, ":")[0]
		appname = strings.Split(appname, ":")[1]
	} else {
		modelname = "base"
	}
	appPath := path.Join(currpath, appname)
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

	goroot := path.Join(build.Default.GOROOT, "./bin/go")

	if err := utils.ExecShell(goroot, "mod", "init", appname); err != nil {
		return fmt.Errorf("go mod init error: %w", err)
	}

	return nil
}

func generateDirAndFile(appdir, modelname string) error {
	appconfig := config.NewAppconfig(modelname)
	if err := appconfig.Parse(); err != nil {
		return fmt.Errorf("%w", err)
	}
	appconfig.Apppath = appdir

	for _, dir := range appconfig.Dirs {
		if err := os.Mkdir(path.Join(appdir, dir), 0755); err != nil {
			return err
		}
		fmt.Printf("mkdir directory %s/%s\n", appdir, dir)
	}

	// create file main.go
	if err := appconfig.GenerateMain(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
