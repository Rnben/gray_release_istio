package grayservice

import (
	"istio.io/api/networking/v1alpha3"

	"log"
)



type InitgrayService struct {
	Gateway               *v1alpha3.Gateway                    `json:"gateway"`
	Destinationrule       *v1alpha3.DestinationRule            `json:"destination"`
	Virtulservice         *v1alpha3.VirtualService             `json:"virtrulservice"`
}

type Initgray struct {
	Virtualservice         *v1alpha3.VirtualService             `json:"virtrulservice"`
	Destinationrule        *v1alpha3.DestinationRule            `json:"destination"`

}

//pre for gray_release

func (svc *InitgrayService) initService(appid,proc,image string,versions []string,rules map[string]int32, gw_port uint32,port int) *InitgrayService {

	return &InitgrayService{
		Gateway:               GetGateway(gw_port,proc),
		Destinationrule:       GetDestinationRule(appid,versions),
		Virtulservice:         GetVs(appid,rules,port),

	}

}


//init_policy of gray_release
func (gp *Initgray) initgray(appid,proc,image string,rules map[string]int32, port int) *Initgray {
	var b_version []string
	for i,_ := range rules{
		b_version = append(b_version,i)
	}
	r := &Initgray{}
	r.Destinationrule = GetDestinationRule(appid,b_version)
	r.Virtualservice =  GetVs(appid,rules,port)
	return  r


}

//update_policy of gray_release
func (gp *Initgray) updatepolicy(appid,proc,image,version string,rules map[string]int32, port int) *Initgray {

	r := &Initgray{}
	r.Virtualservice = GetVs(appid,rules,port)
	return r

}



func  GetGateway(number uint32,pro string) *v1alpha3.Gateway {
	r := &v1alpha3.Gateway{}
	var mapval map[string]string
	mapval["app"] = "ingressgateway"
	r.Selector = mapval
	r.Servers[0].Port.Number = number
	r.Servers[0].Port.Protocol = pro
	return r
}

func GetDestinationRule(appid string, versions []string) *v1alpha3.DestinationRule {
	r := &v1alpha3.DestinationRule{}
	r.Host = appid
	nt := []*v1alpha3.Subset{}
	for _, v := range versions {
		var s *v1alpha3.Subset
		var mapval map[string]string
		mapval["version"] = v
		s.Name = v
		s.Labels = mapval
		nt = append(nt, s)
	}
	r.Subsets = nt

	return r
}

func GetVs(appid string, rules map[string]int32, port int) *v1alpha3.VirtualService {
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
