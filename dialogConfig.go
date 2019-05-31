// Cross-platform.

package cfd

func (config *DialogConfig) apply(dialog Dialog) error {
	var err error
	if config.Role != "" {
		err = dialog.SetTitle(config.Title)
		if err != nil {
			return err
		}
	}
	if config.Role != "" {
		err = dialog.SetRole(config.Role)
		if err != nil {
			return err
		}
	}
	if config.InitialFolder != "" {
		err = dialog.SetInitialFolder(config.InitialFolder)
		if err != nil {
			return err
		}
	}
	if config.DefaultFolder != "" {
		err = dialog.SetDefaultFolder(config.DefaultFolder)
		if err != nil {
			return err
		}
	}
	if config.FileFilter != "" {
		err = dialog.SetFileFilter(config.FileFilter)
		if err != nil {
			return err
		}
	}
	return nil
}
