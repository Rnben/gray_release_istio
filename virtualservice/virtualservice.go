package virtualservice

import (
	"istio.io/api/networking/v1alpha3"
	"log"
)

type VirtualService struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json: "kind"`
	Metadata   Metadata `json: "metadata"`
	Spec       Spec     `json: "spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type Spec struct {
	Hosts    []string `json: "hosts"`
	Http     []Http   `json: "http"`
	Gateways []string `json: "gateways"`
}

type Http struct {
	Route []Route `json:"route"`
	Match []Match `json:"match"`
}

type Match struct {
	Uri Uri `json:"uri"`
}

type Uri struct {
	Exact string `json:"exact"`
}

type Route struct {
	Destination Destination `json:"destination"`
	Weight      int         `json: "weight"`
}

type Destination struct {
	Host   string `json: "host"`
	Port   Port   `json:"port"`
	Subset string `json:"subset"`
}

type Port struct {
	Number int `json:"number"`
}


func (vm *VirtualService) GetVs(appid string, rules map[string]int, port int) *VirtualService {
	vm.Metadata.Name = appid
	vm.Spec.Http[0].Route[0].Destination.Port.Number = port
	vm.Spec.Gateways[0] = appid + "-gateway"
	var rs []Route
	tmp := vm.Spec.Http[0].Route[0]
	tmp.Destination.Host = appid
	for k, v := range rules {
		log.Println(k, v, rs)
		tmp.Destination.Subset = k
		tmp.Weight = v
		rs = append(rs, tmp)
	}
	vm.Spec.Http[0].Route = rs
	return vm
}

func (vm *VirtualService) GetVsCon(appid string, version string) *VirtualService {

	vm.Metadata.Name = appid
	vm.Spec.Gateways[0] = appid + "-gateway"
	vm.Spec.Http[0].Route[0].Destination.Host = appid
	return vm
}


//Todo vir based on offical api
func GetVs2(appid string, rules map[string]int32, port int) *v1alpha3.VirtualService {
	r := &v1alpha3.VirtualService{}
	r.Http[0].Route[0].Destination.Port.Port.Size() = port
	r.Gateways[0] = appid + "-gateway"
	var rs []*v1alpha3.HTTPRouteDestination
	tmp := r.Http[0].Route[0]
	tmp.Destination.Host = appid
	for k, v := range rules {
		log.Println(k, v, rs)
		tmp.Destination.Subset = k
		tmp.Weight = v
		rs = append(rs, tmp)
	}
	r.Http[0].Route = rs
	return r
}