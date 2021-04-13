package shiori

import (
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

func ClearCommand(ctx *cli.Context) error {
	if ctx.NArg() == 1 {
		if !Exists(ctx.Args().Get(0)) {
			return fmt.Errorf("%s is not exists", ctx.Args().Get(0))
		}
		fullPath, _ := filepath.Abs(ctx.Args().Get(0))

		conf, err := os.UserConfigDir()
		if err != nil {
			return err
		}
		err = os.RemoveAll(filepath.Join(conf, "shiori", base64.StdEncoding.EncodeToString([]byte(fullPath))))
		return err
	} else {
		conf, err := os.UserConfigDir()
		if err != nil {
			return err
		}
		err = os.RemoveAll(filepath.Join(conf, "shiori"))
		if err != nil {
			return err
		}
		return makeConfigFile()
	}
}
