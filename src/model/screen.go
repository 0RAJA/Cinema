package model

// Screen 影厅
type Screen struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Rows    int    `json:"rows,omitempty"`
	Cols    int    `json:"cols,omitempty"`
	ImgPath string `json:"img_path,omitempty"`
}
