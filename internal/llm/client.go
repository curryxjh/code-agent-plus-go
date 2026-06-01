package llm

import (
	"code-agent-plus-go/internal/conversation"
	"context"
)

type Client interface {
	Stream(ctx context.Context, conv *conversation.Manager, tools []map[string]any) (<-chan StreamEvent, <-chan error)
}
