package setting

func DoThat(err error, f func() error) error {
	if err != nil {
		return err
	}
	return f()
}
