package main

import (
	"flag"
	"fmt"
	stats "stats-cli/pkg/statistics"
)

type flags struct {
	mean   bool
	mode   bool
	median bool
	sd     bool
}

func main() {
	flags := newFlags()
	flags.parseFlags()
	flags.checkFlags()

	st := stats.NewStat()
	st.Init()
	if flags.mean {
		fmt.Printf("mean: %.2f\n", st.Mean())
	}
	if flags.median {
		fmt.Printf("median: %.2f\n", st.Median())
	}
	if flags.mode {
		fmt.Printf("mode: %d\n", st.Mode())
	}
	if flags.sd {
		fmt.Printf("sd: %.2f\n", st.StandardDeviation())
	}
}

func newFlags() *flags {
	return &flags{mean: false, median: false, mode: false, sd: false}
}

func (f *flags) parseFlags() {
	flag.BoolVar(&f.mean, "mean", false, "Display the arithmetic mean of the sequence")
	flag.BoolVar(&f.mode, "mode", false, "Display the arithmetic mode of the sequence")
	flag.BoolVar(&f.median, "median", false, "Display the arithmetic median of the sequence")
	flag.BoolVar(&f.sd, "sd", false, "Display the arithmetic sd of the sequence")
	flag.Parse()
}

func (f *flags) checkFlags() {
	if !f.mean && !f.mode && !f.median && !f.sd {
		f.setFlagsTrue()
	}
}

func (f *flags) setFlagsTrue() {
	f.mean = true
	f.mode = true
	f.median = true
	f.sd = true
}
