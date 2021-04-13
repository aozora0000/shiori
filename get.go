package shiori

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/peterbourgon/diskv"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func GetCommand(ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return errors.New("Invalid Input File")
	}
	if !Exists(ctx.Args().Get(0)) {
		return fmt.Errorf("%s is not exists", ctx.Args().Get(0))
	}

	fullPath, _ := filepath.Abs(ctx.Args().Get(0))

	conf, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	database := diskv.New(diskv.Options{
		BasePath:     filepath.Join(conf, "shiori"),
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 1024 * 1024,
	})
	return do(fullPath, database)
}

func do(fullPath string, database *diskv.Diskv) error {
	val, err := database.Read(base64.StdEncoding.EncodeToString([]byte(fullPath)))
	if err != nil || string(val) == "" {
		val = []byte("0")
	}

	offsetByte, err := strconv.Atoi(string(val))
	if err != nil {
		return err
	}
	fd, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer fd.Close()

	stat, _ := fd.Stat()
	if int(stat.Size()) < offsetByte {
		offsetByte = int(stat.Size())
	}

	reader := bufio.NewReader(fd)
	reader.Discard(offsetByte)
	writeByte, err := io.Copy(os.Stdout, reader)
	if err != nil {
		return err
	}
	err = database.Write(base64.StdEncoding.EncodeToString([]byte(fullPath)), []byte(strconv.Itoa(int(offsetByte)+int(writeByte))))

	if err != nil {
		return err
	}
	return nil
}
