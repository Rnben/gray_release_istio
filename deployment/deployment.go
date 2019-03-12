package deployment

import "istio.io/api/networking/v1alpha3"
type Deployment struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json: "kind"`
	Metadata   Metadata `json: "metadata"`
	Spec       Spec     `json: "spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type Nmetadata struct {
	Labels NnodeServer `json:"labels"`
}

type NodeServer struct {
	Name    string `json: "name"`
	App     string `json: "app"`
	Version string `json: "version"`
}

type NnodeServer struct {
	App     string `json: "app"`
	Version string `json: "version"`
}

type Spec struct {
	Replicas int       `json: "replicas"`
	Template *Template `json: "template"`
}

type Nspec struct {
	Containers []Container `json: "containers"`
}

type Template struct {
	Metadata Nmetadata `json: "metadata"`
	Spec     Nspec     `json: "spec"`
}

type Container struct {
	Name            string  `json: "name"`
	Image           string  `json: "image"`
	ImagePullPolicy string  `json:"imagePullPolicy"`
	Env             []Env   `json: "env"`
	Ports           []Nport `json: "ports"`
}

type Nport struct {
	ContainerPort int `json:"containerPort"`
}

type Env struct {
	Name  string `json: "name"`
	Value string `json: "value"`
}


func (deploy *Deployment) GetDeploy(appid string, image string, version string, envs []Env, port,num int) *Deployment {
	deploy.Metadata.Name = appid
	deploy.Spec.Replicas = num
	deploy.Spec.Template.Spec.Containers[0].Name = appid
	deploy.Spec.Template.Metadata.Labels.App = appid
	deploy.Spec.Template.Spec.Containers[0].Image = image
	deploy.Metadata.Name = appid + "-" + version
	deploy.Spec.Template.Metadata.Labels.Version = version
	deploy.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort = port

	// 应用基本环境变量
	envs = append(envs, Env{Name: "JAVA_HOME", Value: "/app/jdk"})
	envs = append(envs, Env{Name: "TZ", Value: "Asia/Shanghai"})
	envs = append(envs, Env{Name: "JAVA_OPTS", Value: "-Xms2048m -Xmx2048m -XX:PermSize=128m -XX:MaxPermSize=128m -Dbomc.appname=${APPNAME%.*}"})
	deploy.Spec.Template.Spec.Containers[0].Env = envs

	return deploy
}

//Todo deployment by offical api
func (deploy *v1a) GetDeploy2(appid string, image string, version string, envs []Env, port,num int) *Deployment {
	deploy.Metadata.Name = appid
	deploy.Spec.Replicas = num
	deploy.Spec.Template.Spec.Containers[0].Name = appid
	deploy.Spec.Template.Metadata.Labels.App = appid
	deploy.Spec.Template.Spec.Containers[0].Image = image
	deploy.Metadata.Name = appid + "-" + version
	deploy.Spec.Template.Metadata.Labels.Version = version
	deploy.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort = port

	// 应用基本环境变量
	envs = append(envs, Env{Name: "JAVA_HOME", Value: "/app/jdk"})
	envs = append(envs, Env{Name: "TZ", Value: "Asia/Shanghai"})
	envs = append(envs, Env{Name: "JAVA_OPTS", Value: "-Xms2048m -Xmx2048m -XX:PermSize=128m -XX:MaxPermSize=128m -Dbomc.appname=${APPNAME%.*}"})
	deploy.Spec.Template.Spec.Containers[0].Env = envs

	return deploy
}