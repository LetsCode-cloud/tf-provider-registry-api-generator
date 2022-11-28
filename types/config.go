package types

type Config struct {
	OrganizationName string   `json:"orgName"`
	PluginName       string   `json:"pluginName"`
	Protocols        []string `json:"protocols"`
	RegistryUrl      string   `json:"registryUrl"`
	PgpKeyId         string   `json:"pgpKeyId"`
}
