package appoptions

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.MySQLOptions.Validate()...)

	return errs
}