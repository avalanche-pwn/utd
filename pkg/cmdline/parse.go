package cmdline

import (
	"flag"
)

type Flags struct {
	Port  int
	Host  string
	Debug bool //true if exists
}

func (f *Flags) ParseFlags() {
	flag.IntVar(&f.Port, "p", 1337, "port?")
	flag.StringVar(&f.Host, "h", "127.0.0.1", "host?")
	flag.BoolVar(&f.Debug, "d", false, "")

	flag.Parse()
	// fmt.Print(f.Debug)
}
