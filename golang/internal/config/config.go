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
	Port string `default:":8000"`
}
