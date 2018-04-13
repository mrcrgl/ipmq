package cli

import "flag"

func FlagSet(options *Options) *flag.FlagSet {
	fs := flag.NewFlagSet(Name, flag.ExitOnError)

	setFlags(fs, options)

	return fs
}

func setFlags(set *flag.FlagSet, options *Options) {
	set.BoolVar(&options.Debug, "debug", options.Debug, "increase verbosity")
}
