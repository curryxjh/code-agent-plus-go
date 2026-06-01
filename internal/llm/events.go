package llm

type StreamEvent interface {
	streamEvent()
}

type TextDelta struct { // 文本片段
	Text string
}
type ThinkingDelta struct { // 思考过程片段
	Text string
}

type ThinkingComplete struct { // 思考完成
	Thinking  string
	Signature string
}
type ToolCallStart struct { // 工具调用开始
	ToolName, ToolID string
}

type ToolCallDelta struct { // 工具参数片段
	Text string
}

type ToolCallComplete struct { // 工具调用完成
	ToolID    string
	ToolName  string
	Arguments map[string]any
}
type UsageInfo struct {
	InputTokens  int
	OutputTokens int
}

type StreamEnd struct { // 流结束
	StopReason string
	Usage      UsageInfo
}
