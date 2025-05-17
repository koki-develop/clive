package cli

import (
	_ "embed"
	"fmt"

	"github.com/koki-develop/clive/internal/util"
)

//go:embed static/clive.yml
var configInitTemplate []byte

type InitParams struct {
	Config string
}

func (c *CLI) Init(p *InitParams) error {
	exists, err := util.FileExists(p.Config)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("%s already exists", p.Config)
	}

	f, err := util.CreateFile(p.Config)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(configInitTemplate); err != nil {
		return err
	}

	if _, err = fmt.Fprintf(c.stdout, "Created %s\n", p.Config); err != nil {
		return err
	}

	return nil
}
