package gocd

const (
	AgentsEndpoint           = "/api/agents"
	VersionEndpoint          = "/api/version"
	ServerHealthEndpoint     = "/api/server_health_messages"
	ConfigReposEndpoint      = "/api/admin/config_repos"
	SystemAdminEndpoint      = "/api/admin/security/system_admins"
	BackupConfigEndpoint     = "/api/config/backup"
	PipelineGroupEndpoint    = "/api/admin/pipeline_groups"
	EnvironmentEndpoint      = "/api/admin/environments"
	JobRunHistoryEndpoint    = "/api/agents/%s/job_run_history"
	MaintenanceEndpoint      = "/api/admin/maintenance_mode"
	APIFeedPipelineEndpoint  = "/api/feed/pipelines.xml"
	PipelineStatus           = "/api/pipelines/%s/status"
	EncryptEndpoint          = "/api/admin/encrypt"
	ArtifactInfoEndpoint     = "/api/admin/config/server/artifact_config"
	PipelinesEndpoint        = "/api/pipelines"
	HealthEndpoint           = "/api/v1/health"
	DefaultTimeoutEndpoint   = "api/admin/config/server/default_job_timeout"
	MailServerConfigEndpoint = "api/config/mailserver"
	PluginSettingsEndpoint   = "/api/admin/plugin_settings"
	AuthConfigEndpoint       = "/api/admin/security/auth_configs"
	SiteURLEndpoint          = "/api/admin/security/site_urls"
	HeaderVersionOne         = "application/vnd.go.cd.v1+json"
	HeaderVersionTwo         = "application/vnd.go.cd.v2+json"
	HeaderVersionThree       = "application/vnd.go.cd.v3+json"
	HeaderVersionFour        = "application/vnd.go.cd.v4+json"
	HeaderVersionSeven       = "application/vnd.go.cd.v7+json"
)

const (
	goCdAPILoggerName = "gocd-sdk-go"
	ContentJSON       = "application/json"
	HeaderConfirm     = "X-GoCD-Confirm"
)
