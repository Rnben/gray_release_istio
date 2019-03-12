package service

type Service struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json: "kind"`
	Metadata   Metadata `json: "metadata"`
	Spec       Spec     `json: "spec"`
}

type Metadata struct {
	Name   string     `json:"name"`
	Labels NodeServer `json:"labels"`
}

type NodeServer struct {
	App string `json: "app"`
}

type Spec struct {
	// Ports    []Port   `json: "containers"`
	Ports    []Port   `json: "ports"`
	Selector Selector `json:"selector"`
}

type Selector struct {
	App string `json:"app"`
}

type Port struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type Env struct {
	Name  string `json: "name"`
	Value string `json: "value"`
}


func (service *Service) GetService(appid string, port int) *Service {
	service.Metadata.Name = appid
	service.Metadata.Labels.App = appid
	service.Spec.Selector.App = appid
	service.Spec.Ports[0].Port = port

	return service
}
