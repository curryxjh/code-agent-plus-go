package llm

import (
	"code-agent-plus-go/internal/config"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/anthropics/anthropic-sdk-go/packages/param"
	"strings"
	"time"
)

const anthropicStreamIdleTimeout = 5 * time.Minute

type anthropicClient struct {
	client          anthropic.Client
	model           string
	thinking        bool
	systemPrompt    string
	maxOutputTokens int
	contextWindow   int
}

func supportAdaptiveThinking(model string) bool {
	for _, family := range []string{"claude-opus-4-", "claude-sonnet-4-"} {
		if strings.HasPrefix(model, family) {
			rest := model[len(family):]
			if len(rest) > 0 && rest[0] >= '6' && rest[0] <= '9' {
				return true
			}
		}
	}
	return false
}

func newAnthropicClient(cfg *config.ProviderConfig, systemPrompt string) (*anthropicClient, error) {
	apiKey := cfg.ResolveAPIKey()
}
