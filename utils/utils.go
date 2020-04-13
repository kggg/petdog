package utils
import (
	"os"
	"os/exec"
)

// IsExist check file or directory exists in the path or not, return true or false
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func ExecShell(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	if err := cmd.Run(); err != nil{
		return err
	}
	return nil
}