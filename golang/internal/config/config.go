package config

// AppConfig holds the application config
// type AppConfig struct {
// 	UseCache      bool
// 	TemplateCache map[string]*template.Template
// 	InfoLog       *log.Logger
// 	ErrorLog      *log.Logger
// 	InProduction  bool
// 	Session       *scs.SessionManager
// }

type AppConfig struct {
	Port      string `default:":8000"`
	BuildInfo string `default:"Version: 1.0.0, BuildTime: 1970-01-01T00:00:00Z, GitCommit: Testing"`
}

func (c *AppConfig) GetBuildInfo() string {
	return c.BuildInfo
}
