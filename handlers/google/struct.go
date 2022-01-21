package google

type NotificationStruct struct {
	Incident Incident `json:"incident,omitempty"`
	Version  string   `json:"version,omitempty"`
}

type Incident struct {
	ScopingProjectID        interface{} `json:"scoping_project_id,omitempty"`         // 托管指标范围的项目 ID
	URL                     string      `json:"url,omitempty"`                        // 突发事件的 Google Cloud Console 网址
	StartedAt               int64       `json:"started_at,omitempty"`                 // 开始时间
	EndedAt                 int64       `json:"ended_at,omitempty"`                   // 结束时间
	ResourceDisplayName     string      `json:"resource_display_name,omitempty"`      // 资源展示名称
	ResourceTypeDisplayName string      `json:"resource_type_display_name,omitempty"` // 监控资源类型的显示名
	PolicyName              string      `json:"policy_name,omitempty"`                // 策略名称
	ConditionName           string      `json:"condition_name,omitempty"`             // 条件的显示名
	ThresholdValue          string      `json:"threshold_value,omitempty"`            // 阈值
	ObservedValue           string      `json:"observed_value,omitempty"`             // 触发值
	Summary                 string      `json:"summary,omitempty"`
}
