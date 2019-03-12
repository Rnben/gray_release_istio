package grayservice

import (
	"deploy_app/gateway"
	"deploy_app/destination"
	"deploy_app/virtualservice"
	"deploy_app/deployment"
)



type Service struct {
	Gateway               *gateway.GateWay                        `json:"gateway"`
	Destinationrule       *destination.DestinationRule            `json:"destination"`
	Virtulservice         *virtualservice.VirtualService          `json:"virtrulservice"`
}

type Initgray struct {
	*Graypolicy
	Destinationrule       *destination.DestinationRule            `json:"destination"`

}

type Graypolicy struct {
	Deployment            deployment.Deployment                  `json:"deployment"`

	Virtualservice        virtualservice.VirtualService          `json:"virtrulservice"`

}

func (svc *Service) initService(appid,proc,image string,versions []string,rules map[string]int, envs []deployment.Env, gw_port,port int) *Service {

	return &Service{
		Gateway:               svc.Gateway.GetGateway(appid,gw_port,proc),
		Destinationrule:       svc.Destinationrule.GetDestinationRule(appid,versions),
		Virtulservice:         svc.Virtulservice.GetVs(appid,rules,port),

	}

}


//新版本version,各版本及对应值流量权重组成的rules
func (gp *Initgray) initgray(appid,proc,image,version string,rules map[string]int, envs []deployment.Env, port int) *Initgray {
	var b_version []string
	for i,_ := range rules{
		b_version = append(b_version,i)
	}
	r := &Initgray{}
	r.Deployment = *gp.Deployment.GetDeploy(appid,image,version,envs,port,0)
	r.Destinationrule = *gp.Destinationrule.GetDestinationRule(appid,b_version)
	r.Virtualservice = *gp.Virtualservice.GetVs(appid,rules,port)
	return  r


}


func (gp *Initgray) updatepolicy(appid,proc,image,version string,rules map[string]int, envs []deployment.Env, port,replices int) *Graypolicy {

	r := &Graypolicy{}
	r.Virtualservice = *gp.Virtualservice.GetVs(appid,rules,port)
	r.Deployment =  *gp.Deployment.GetDeploy(appid,image,version,envs,port,replices)
	return r

}



//Todo initgray based on offical  api

func (svc *Service) initService2(appid,proc,image string,versions []string,rules map[string]int, envs []deployment.Env, gw_port,port int) *Service {

	return &Service{
		Gateway:               svc.Gateway.GetGateway(appid,gw_port,proc),
		Destinationrule:       svc.Destinationrule.GetDestinationRule(appid,versions),
		Virtulservice:         svc.Virtulservice.GetVs(appid,rules,port),

	}

}


//新版本version,各版本及对应值流量权重组成的rules
func (gp *Initgray) initgray2(appid,proc,image,version string,rules map[string]int, envs []deployment.Env, port int) *Initgray {
	var b_version []string
	for i,_ := range rules{
		b_version = append(b_version,i)
	}
	r := &Initgray{}
	r.Deployment = *gp.Deployment.GetDeploy(appid,image,version,envs,port,0)
	r.Destinationrule = *gp.Destinationrule.GetDestinationRule(appid,b_version)
	r.Virtualservice = *gp.Virtualservice.GetVs(appid,rules,port)
	return  r


}


func (gp *Initgray) updatepolicy2(appid,proc,image,version string,rules map[string]int, envs []deployment.Env, port,replices int) *Graypolicy {

	r := &Graypolicy{}
	r.Virtualservice = *gp.Virtualservice.GetVs(appid,rules,port)
	r.Deployment =  *gp.Deployment.GetDeploy(appid,image,version,envs,port,replices)
	return r

}