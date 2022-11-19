package model

type OS struct {
	Type    string `json:"type,omitempty"`
	Windows bool   `json:"windows"`
	Linux   bool   `json:"linux"`
	Darwin  bool   `json:"darwin"`
}

type Product struct {
	EdgeVersion  string `json:"edge_version"`
	FlowVersion  string `json:"flow_version"`
	ImageVersion string `json:"image_version"`
	Product      string `json:"product"` // RubixCompute, RubixComputeIO, RubixCompute5, Server, Edge28, Nuc  see https://github.com/NubeIO/lib-command/blob/master/product/product.go#L7
	Arch         string `json:"arch"`    // armv7 amd64
	OS           OS     `json:"os"`      // Linux, Windows, Darwin
}
