package domain

// AgentTask does not contain provider-specific process or protocol details.
type AgentTask struct {
	Project string
	Prompt  string
	Model   string
}
type ProviderStatus struct {
	Available    bool   `json:"available"`
	LoginPresent bool   `json:"login_present"`
	Version      string `json:"version,omitempty"`
	Detail       string `json:"detail,omitempty"`
}
