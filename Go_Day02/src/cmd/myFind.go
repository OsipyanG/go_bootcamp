package cmd

import (
	"errors"
	"flag"
	"log"
)

type CliFlags struct {
	sl       bool
	d        bool
	f        bool
	ext      string
	filePath string
}

var ErrExtFlag = errors.New("the --ext flag can only work in combination with the -f flag")
var ErrPathNotSet = errors.New("the search path is not set")

func (c *CliFlags) initFlags() error {
	flag.BoolVar(&c.sl, "sl", false, "sl")
	flag.BoolVar(&c.d, "d", false, "d")
	flag.BoolVar(&c.f, "f", false, "f")
	flag.StringVar(&c.ext, "ext", "", "ext")
	flag.StringVar(&c.filePath, "path", "", "path")
	flag.Parse()
	if c.ext != "" && !c.f {
		return ErrExtFlag
	}
	if c.filePath == "" {
		return ErrPathNotSet
	}
	c.fillFlags()
	return nil
}

func (c *CliFlags) fillFlags() {
	if !c.sl && !c.d && !c.f {
		c.sl = true
		c.d = true
		c.f = true
	}
}

func main() {
	cliFlags := CliFlags{}
	err := cliFlags.initFlags()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
