package hooks

type Input[TI any, TR any] struct {
	Event        string `json:"hook_event_name"`
	Tool         string `json:"tool_name"`
	ToolInput    TI     `json:"tool_input"`
	ToolResponse TR     `json:"tool_response"`
}

type FileInput struct {
	Path string `json:"file_path"`
}
