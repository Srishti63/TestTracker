package dto

type PresetConfigItem struct {
	SubjectID  string  `json:"subject_id" binding:"required,uuid4"`
	TotalMarks float64 `json:"total_marks" binding:"required,gt=0"`
}

type BatchPresetRequest struct {
	TestGroupID string             `json:"test_group_id" binding:"required,uuid4"`
	Presets     []PresetConfigItem `json:"presets" binding:"required,dive"`
}

type PresetConfigResponse struct {
	ID          string  `json:"id"`
	TestGroupID string  `json:"test_group_id"`
	SubjectID   string  `json:"subject_id"`
	SubjectName string  `json:"subject_name"`
	TotalMarks  float64 `json:"total_marks"`
}