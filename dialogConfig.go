// Cross-platform.

package cfd

var defaultFilters = []FileFilter{
	{
		DisplayName: "All Files (*.*)",
		Pattern:     "*.*",
	},
}

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
	var fileFilters []FileFilter

	if config.FileFilters != nil && len(config.FileFilters) > 0 {
		fileFilters = config.FileFilters
	} else {
		fileFilters = defaultFilters
	}
	err = dialog.SetFileFilter(fileFilters)
	if err != nil {
		return err
	}
	return nil
}
