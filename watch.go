package shiori

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/peterbourgon/diskv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
)

func WatchCommand(ctx *cli.Context) error {
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
	err = do(fullPath, database)
	if err != nil {
		return err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan error)
	go func() {
		for {
			select {
			case event := <-watcher.Events:

				if event.Op&fsnotify.Write == fsnotify.Write {
					err := do(fullPath, database)
					if err != nil {
						done <- err
					}
				}
			case err := <-watcher.Errors:
				log.Println("done")
				done <- err
			}
		}
	}()
	err = watcher.Add(fullPath)
	if err != nil {
		return err
	}
	return <-done
}
