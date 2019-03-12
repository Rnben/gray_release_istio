package destination

import "istio.io/api/networking/v1alpha3"

type DestinationRule struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec`
}

type Metadata struct {
	Name string `json: "name"`
}

type Spec struct {
	Host    string   `json:"host"`
	Subsets []Subset `json:""subsets`
}

type Subset struct {
	Name   string `json: "name"`
	Labels Node   `json: "labels"`
}

type Node struct {
	Version string `json:"version"`
}


func (rule *DestinationRule) GetDestinationRule(appid string, versions []string) (*DestinationRule) {



	rule.Metadata.Name = appid
	rule.Spec.Host = appid
	nt := rule.Spec.Subsets
	for _, v := range versions {
		var s Subset
		s.Name = v
		s.Labels.Version = v
		nt = append(nt, s)
	}
	rule.Spec.Subsets = nt
	apiver := "networking.istio.io/v1alpha3"
	kindval := "DestinationRule"
	return &DestinationRule{
		ApiVersion:   apiver,
		Kind:		  kindval,
		Metadata:     Metadata{appid},
		Spec:         *getspec(appid,versions),
	}
}



func getspec(appid string, versions []string) *Spec{
	subsetval := []Subset{}
	for _,ival := range versions{

		verval := ival
		subsetval = append(subsetval,*getsubset(verval))
	}
	return &Spec{
		Host:         appid,
		Subsets:      subsetval,
	}

}

func getsubset(version string)  *Subset{


	return &Subset{
		Name :		version,
		Labels:     Node{version},
	}

}

//Todo func based on offical api
func GetDestinationRule2(appid string, versions []string) *v1alpha3.DestinationRule {
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