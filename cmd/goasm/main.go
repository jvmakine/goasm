package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/jvmakine/goasm/class"
	"github.com/jvmakine/goasm/classfile"
	"gopkg.in/yaml.v3"
)

var CLI struct {
	CatCmd `cmd:""`
}

type CatCmd struct {
	ClassFile string `arg:"" help:"the .class file to show"`
}

func (r *CatCmd) Run() error {
	file, err := os.Open(r.ClassFile)
	if err != nil {
		return err
	}
	classFile, err := classfile.Parse(file)
	if err != nil {
		return err
	}
	clazz := class.NewClass(classFile)
	summary := SummaryFrom(clazz, classFile)
	bytes, err := yaml.Marshal(summary)
	if err != nil {
		return err
	}
	println(string(bytes))
	return nil
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
