package conf

type Options struct {
	ListenIP    string      `yaml:"ListenIP"`
	SendIP      string      `yaml:"SendIP"`
	LimitConfig LimitConfig `yaml:"LimitConfig"`
	CertConfig  *CertConfig `yaml:"CertConfig"`
	XrayOptions XrayOptions `yaml:"XrayOptions"`
	HyOptions   HyOptions   `yaml:"HyOptions"`
	SingOptions SingOptions `yaml:"SingOptions"`
}

type XrayOptions struct {
	EnableProxyProtocol bool                    `yaml:"EnableProxyProtocol"`
	EnableDNS           bool                    `yaml:"EnableDNS"`
	DNSType             string                  `yaml:"DNSType"`
	EnableUot           bool                    `yaml:"EnableUot"`
	EnableTFO           bool                    `yaml:"EnableTFO"`
	DisableIVCheck      bool                    `yaml:"DisableIVCheck"`
	DisableSniffing     bool                    `yaml:"DisableSniffing"`
	EnableFallback      bool                    `yaml:"EnableFallback"`
	FallBackConfigs     []FallBackConfigForXray `yaml:"FallBackConfigs"`
}

type FallBackConfigForXray struct {
	// xray sing-box
	Alpn string `yaml:"Alpn"`
	// xray
	SNI              string `yaml:"SNI"`
	Path             string `yaml:"Path"`
	Dest             string `yaml:"Dest"`
	ProxyProtocolVer uint64 `yaml:"ProxyProtocolVer"`
	// sing-box
	Server     string `yaml:"Server"`
	ServerPort string `yaml:"ServerPort"`
}

type FallBackConfigForSing struct {
	// sing-box
	FallBack        FallBack            `yaml:"FallBack"`
	FallBackForALPN map[string]FallBack `yaml:"FallBackForALPN"`
}
type FallBack struct {
	Server     string `yaml:"Server"`
	ServerPort string `yaml:"ServerPort"`
}

type HyOptions struct {
	Resolver          string `yaml:"Resolver"`
	ResolvePreference string `yaml:"ResolvePreference"`
	SendDevice        string `yaml:"SendDevice"`
}

type SingOptions struct {
	EnableProxyProtocol      bool                  `yaml:"EnableProxyProtocol"`
	EnableDNS                bool                  `yaml:"EnableDNS"`
	EnableTUIC               bool                  `yaml:"EnableTUIC"`
	CongestionControl        string                `yaml:"CongestionControl"`
	DomainStrategy           string                `yaml:"DomainStrategy"`
	TCPFastOpen              bool                  `yaml:"EnableTFO"`
	SniffEnabled             bool                  `yaml:"EnableSniff"`
	SniffOverrideDestination bool                  `yaml:"SniffOverrideDestination"`
	FallBackConfigs          FallBackConfigForSing `yaml:"FallBackConfigs"`
}
