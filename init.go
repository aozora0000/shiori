package shiori

import (
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

func makeConfigFile() error {
	conf, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(conf, "shiori"), 0700)
	if err != nil {
		return err
	}
	return nil
}

func InitCommand(_ *cli.Context) error {
	return makeConfigFile()
}
