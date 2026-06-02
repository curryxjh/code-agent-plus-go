package llm

import (
	"code-agent-plus-go/internal/config"
	"code-agent-plus-go/internal/conversation"
	"context"
)

type Client interface {
	Stream(ctx context.Context, conv *conversation.Manager, tools []map[string]any) (<-chan StreamEvent, <-chan error)
}

type MaxTokenSetter interface {
	SetMaxOutputTokens(tokens int)
}

func NewClient(cfg *config.ProviderConfig, systemPrompt string) (Client, error) {
	switch cfg.Protocol {
	case "anthropic":

	}
}
