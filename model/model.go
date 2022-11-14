package model

type UIElement struct {
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	Message string `json:"message"`
	UIType  string `json:"ui_type"`
}
