package cli

type Options struct {
	Debug bool
}

func (o *Options) Validate() error {
	/*if o.Debug == false {
		return errors.New("debug needs to be set")
	}*/

	return nil
}

func DefaultOptions() *Options {
	options := new(Options)

	options.Debug = false

	return options
}
