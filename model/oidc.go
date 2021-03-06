package model

type Oidc struct {
	Enabled         bool   `yaml:"enabled"`
	IssuerUrl       string `yaml:"issuerUrl"`
	ClientId        string `yaml:"clientId"`
	UsernameClaim   string `yaml:"usernameClaim"`
	GroupsClaim     string `yaml:"groupsClaim,omitempty"`
	UsernamePrefix  string `yaml:"usernamePrefix,omitempty"`
	GroupsPrefix    string `yaml:"groupsPrefix,omitempty"`
}
