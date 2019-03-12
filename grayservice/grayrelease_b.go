package grayservice

import (
	"istio.io/api/networking/v1alpha3"


	"deploy_app/deployment"
	"deploy_app/destination"
	"deploy_app/virtualservice"

	"deploy_app/gateway"
)



type InitgrayService2 struct {
	Gateway               *v1alpha3.Gateway                    `json:"gateway"`
	Destinationrule       *v1alpha3.DestinationRule            `json:"destination"`
	Virtulservice         *v1alpha3.VirtualService             `json:"virtrulservice"`
}

type Initgray2 struct {
	*Graypolicy
	Destinationrule        *v1alpha3.DestinationRule            `json:"destination"`

}

type Graypolicy2 struct {
	//Deployment                 Todo            `json:"deployment"`

	Virtualservice        *v1alpha3.VirtualService          `json:"virtrulservice"`

}

func (svc *InitgrayService2) initService2(appid,proc,image string,versions []string,rules map[string]int32, envs []deployment.Env, gw_port uint32,port int) *InitgrayService2 {

	return &InitgrayService2{
		Gateway:               gateway.GetGateway2(appid,gw_port,proc),
		Destinationrule:       destination.GetDestinationRule2(appid,versions),
		Virtulservice:         virtualservice.GetVs2(appid,rules,port),

	}

}


//新版本version,各版本及对应值流量权重组成的rules
func (gp *Initgray2) initgray2(appid,proc,image,version string,rules map[string]int, envs []deployment.Env, port int) *Initgray2 {
	var b_version []string
	for i,_ := range rules{
		b_version = append(b_version,i)
	}
	r := &Initgray2{}
	r.Deployment = *gp.Deployment.GetDeploy(appid,image,version,envs,port,0)
	r.Destinationrule = destination.GetDestinationRule2(appid,b_version)
	r.Virtualservice = *gp.Virtualservice.GetVs(appid,rules,port)
	return  r


}


func (gp *Initgray2) updatepolicy(appid,proc,image,version string,rules map[string]int, envs []deployment.Env, port,replices int) *Initgray2 {

	r := &Initgray2{}
	r.Virtualservice = *gp.Virtualservice.GetVs(appid,rules,port)
	r.Deployment =  *gp.Deployment.GetDeploy(appid,image,version,envs,port,replices)
	return r

}


