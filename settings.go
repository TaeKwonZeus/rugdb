package rugdb

type Settings struct {
	PageSize int
}

func fillDefaults(settings Settings) Settings {
	if settings.PageSize == 0 {
		settings.PageSize = defaultPageSize
	}

	return settings
}

var (
	defaultPageSize = 0x1000
)
