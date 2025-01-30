package config

type CloudConfig struct {
	SelectedCloud string
	AWSConfig
	GCPConfig
	OSTConfig
}

type AWSConfig struct {
	Profile        string
	Region         string
	UseLocalStack  bool
	LocalStackPort string
}
type GCPConfig struct {
	CredFilePath string
}

type OSTConfig struct {
	CloudsFilePath string
}

func NewCloudConfig() CloudConfig {

	return CloudConfig{}
}
