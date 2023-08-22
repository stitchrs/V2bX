package conf

type SingConfig struct {
	LogConfig     SingLogConfig `yaml:"Log"`
	OriginalPath  string        `yaml:"OriginalPath"`
	DnsConfigPath string        `yaml:"DnsConfigPath"`
	NtpConfig     SingNtpConfig `yaml:"NTP"`
}

type SingLogConfig struct {
	Disabled  bool   `yaml:"Disable"`
	Level     string `yaml:"Level"`
	Output    string `yaml:"Output"`
	Timestamp bool   `yaml:"Timestamp"`
}

type SingNtpConfig struct {
	Enable bool   `yaml:"Enable"`
	Server string `yaml:"Server"`
}

func NewSingConfig() *SingConfig {
	return &SingConfig{
		LogConfig: SingLogConfig{
			Level:     "error",
			Timestamp: true,
		},
	}
}
