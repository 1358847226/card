package models

type Role struct {
	Id int `json:"id"`
	Parent_id int `json:"parent_id"`
	Name	string `json:"name"`
	Path    string `json:"path"`
	Meta_id int `json:"meta_id"`
	Redirect string `json:"redirect"`
	Component string `json:"component"`
}
