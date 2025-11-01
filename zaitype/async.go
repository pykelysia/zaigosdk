package zaitype

type (
	AsyncResponse struct {
		Model      string `json:"model"`
		ID         string `json:"id"`
		RequestID  string `json:"request_id"`
		TaskStatus string `json:"task_status"`
		Error      `json:"error"`
	}
	GetResponse struct {
		ID            string          `json:"id"`
		RequestID     string          `json:"request_id"`
		Model         string          `json:"model"`
		Choices       []Choices       `json:"choices"`
		VideoResult   []VideoResult   `json:"video_result"`
		WebSearch     []WebSearch     `json:"web_search"`
		ContentFilter []ContentFilter `json:"content_filter"`
		Usage         `json:"usage"`
		Error         `json:"error"`
	}
)
