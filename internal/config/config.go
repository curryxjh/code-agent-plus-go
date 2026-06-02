package config

import (
	"os"
	"strings"
)

var envKeyMap = map[string]string{
	"anthropic":     "ANTHROPIC_API_KEY",
	"openai":        "OPENAI_API_KEY",
	"openai-compat": "OPENAI_API_KEY",
}

var validProtocols = map[string]bool{
	"anthropic":     true,
	"openai":        true,
	"openai-compat": true,
}

type ConfigError struct {
	Message string
}

func (e *ConfigError) Error() string {
	return e.Message
}

type ProviderConfig struct {
	Name            string `yaml:"name"`
	Protocol        string `yaml:"protocol"`
	BaseURL         string `yaml:"base_url"`
	Model           string `yaml:"model"`
	APIKey          string `yaml:"api_key"`
	Thinking        bool   `yaml:"thinking"`
	ContextWindow   int    `yaml:"context_window"`
	MaxOutputTokens int    `yaml:"max_output_tokens"`
}

func (p *ProviderConfig) GetContextWindow() int {
	if p.ContextWindow > 0 {
		return p.ContextWindow
	}
	if strings.Contains(p.Model, "claude") {
		return 200000
	}
	return 128000
}

func (p *ProviderConfig) GetMaxOutputTokens() int {
	if p.MaxOutputTokens > 0 {
		return p.MaxOutputTokens
	}
	if p.Thinking {
		return 64000
	}
	return 8192
}

func (p *ProviderConfig) ResolveAPIKey() string {
	if p.APIKey != "" {
		return p.APIKey
	}
	envVar := envKeyMap[p.Protocol]
	if envVar == "" {
		return ""
	}
	return os.Getenv(envVar)
}

type MCPServerConfig struct {
	Name      string            `yaml:"name"`
	Command   string            `yaml:"command"`
	Args      []string          `yaml:"args"`
	URL       string            `yaml:"url"`
	Transport string            `yaml:"transport"`
	Headers   map[string]string `yaml:"headers"`
	Env       map[string]string `yaml:"env"`
}

type AppConfig struct {
	Providers      []ProviderConfig  `yaml:"providers"`
	PermissionMode string            `yaml:"permission_mode"`
	MCPServers     []MCPServerConfig `yaml:"mcp_servers"`
	Hooks          []hooks.Hooks     `yaml:"hooks"`
}

func loadSingleFile(path string) (*AppConfig, error) {

}
