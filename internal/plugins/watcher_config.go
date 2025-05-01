package plugins

type WatcherConfig struct {
	ID                  string     `json:"id"` // Unique watcher ID
	WatchDir            string     `json:"watch_dir"`
	FileTemplate        string     `json:"file_template"`
	PostURL             string     `json:"post_url"`
	PluginPath          string     `json:"plugin_path"`
	ProcessorPluginPath string     `json:"processor_plugin_path"`
	PostPluginPath      string     `json:"post_plugin_path"`
	AuthPluginPath      string     `json:"auth_plugin_path"`
	ProcessedDir        string     `json:"processed_dir"`
	Auth                AuthConfig `json:"auth"`
}

type AuthConfig struct {
	Cache          string `json:"cache"`
	URL            string `json:"url"`
	BodyTemplate   string `json:"body_template"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	TokenJSONField string `json:"token_jsonfield"`
}
