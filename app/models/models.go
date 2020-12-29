package models

// Datas array of NodeData
type Datas struct {
	Datas []NodeData `json:"data"`
}

// NodeData data in a node
type NodeData struct {
	Data int `json:"value"`
}
