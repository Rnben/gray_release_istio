package main

import (
	"deploy_app/deployment"
	"deploy_app/destination"
	"deploy_app/gateway"
	"deploy_app/service"
	"deploy_app/virtualservice"
	"flag"
	"log"
	"regexp"

	"fmt"
	"os"
	"os/exec"
)

var (
	appid     string
	image     string
	versions  []string
	port      int
	version   string
	mode      string
	namespace string
)

func compressStr(str string) string {
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

func init() {
	flag.StringVar(&namespace, "namespace", "", "app's namespace")
	flag.StringVar(&mode, "mode", "test", "mode：normal、canary、test、ajust")
	flag.StringVar(&appid, "appid", "", "app's name")
	flag.StringVar(&image, "image", "test.registry.zj.chinamobile.com/special/demo-go:latest", "app's image")
	flag.IntVar(&port, "port", 8001, "app's port")
	flag.Parse()
}

func check() {
	if mode != "normal" && mode != "canary" && mode != "test" && mode == "just" {
		fmt.Println("Please input mode")
		os.Exit(1)
	}

	if appid == "" {
		fmt.Println("Please input appid and image")
		os.Exit(1)
	}
	fmt.Println("-----------------")
	if compressStr(namespace) == "" {
		fmt.Println(appid, namespace, image, port, "default")

	} else {
		fmt.Println(appid, namespace, image, port, namespace)
	}
}

func initDeploy(appid string, version string, image string) {
	var d deployment.Deployment
	var envs []deployment.Env
	var key string
	var value string
	for {
		fmt.Println(version, "Please enter app's env like: key value, d to Exit")
		fmt.Scanln(&key, &value)
		if key == "d" {
			break
		}
		e := deployment.Env{Name: key, Value: value}
		envs = append(envs, e)
	}
	log.Println("envs:", envs)
	d.GetDeploy(appid, image, version, envs, port)
}

func initService(appid string, port int) {
	var s service.Service
	s.GetService(appid, port)
}

func initRules(appid string, versions []string) {
	var r destination.DestinationRule
	r.GetDestinationRule(appid, versions)
}

func initVirtualService(appid,version string,port int) {
	var v virtualservice.VirtualService
	valmap :=make(map[string]int)
	valmap[version] = 100
	_ = v.GetVs(appid, valmap,port)
}

func initGateway(appid,pro string, number int) {
	var gateway gateway.GateWay
	gateway.GetGateway(appid, number, pro)
}

func updateVirtualService(appid string, versions map[string]int, port int) {
	var v virtualservice.VirtualService
	v.GetVs(appid, versions, port)
}

func run_command(s string) string {
	var output []byte
	var cmd *exec.Cmd
	var err error
	cmd = exec.Command("bash", "-c", s)
	if output, err = cmd.Output(); err != nil {
		log.Fatal(err)
	}
	return string(output)
}

func normal(namespace string) {
	log.Println("---------start-----------")
	log.Println("Please enter app's version")
	fmt.Scanln(&version)
	version = compressStr(version)
	if version == "" {
		log.Fatal("version vaild")
	}
	initDeploy(appid, version, image)
	versions = append(versions, version)
	initService(appid, port)

	var can string
	fmt.Println("Deploy ", appid, "'s ", version, "version ?")
	fmt.Scanln(&can)
	if can == "y" || can == "Y" {
		log.Println("kubectl apply -f ./tmp/")
		if compressStr(namespace) != "" {
			output := run_command("kubectl apply -f <(istioctl kube-inject -f ./tmp/deploy-" + appid + "-" + version + ".json -n " + compressStr(namespace) + ")")
			log.Println(output)
			output = run_command("kubectl apply -f ./tmp/svc-" + appid + ".json -n " + compressStr(namespace))
			log.Println(output)
		} else {
			output := run_command("kubectl apply -f <(istioctl kube-inject -f ./tmp/deploy-" + appid + "-" + version + ".json)")
			log.Println(output)
			output = run_command("kubectl apply -f ./tmp/svc-" + appid + ".json")
			log.Println(output)
		}

	}
}

func canary(namespace string) {
	if len(versions) == 0 {
		var tmp string
		fmt.Println("Please enter canary app's current version")
		fmt.Scanln(&tmp)
		versions = append(versions, tmp)
	}
	// 灰度版本
	log.Println("canary version test")
	fmt.Println("Please enter canary app's version")
	fmt.Scanln(&version)
	version = compressStr(version)
	if version == "" {
		log.Fatal("version vaild")
	}
	versions = append(versions, version)

	initRules(appid, versions)
	initVirtualService(appid, versions[0],port)
	initGateway(appid, "http",port)
	initDeploy(appid, version, image)
	can := ""
	fmt.Println("Deploy ", appid, "'s ", version, "canary version ?")
	fmt.Scanln(&can)

	if can == "y" || can == "Y" {
		log.Println("create gateway-bomc-test")
		if compressStr(namespace) != "" {
			output := run_command("kubectl apply -f  ./tmp/gateway-" + appid + ".json -n " + compressStr(namespace))
			log.Println("create destinationrules")
			output = run_command("kubectl apply -f ./tmp/rules-" + appid + ".json -n " + compressStr(namespace))
			log.Println("create virtualservice,defalut v1 100%")
			output = run_command("kubectl apply -f ./tmp/vm-" + appid + ".json -n " + compressStr(namespace))
			log.Println(output)
		} else {
			output := run_command("kubectl apply -f ./tmp/gateway-" + appid + ".json")
			log.Println("create destinationrules")
			log.Println("create destinationrules")
			output = run_command("kubectl apply -f ./tmp/rules-" + appid + ".json")
			log.Println("create virtualservice,defalut v1 100%")
			output = run_command("kubectl apply -f ./tmp/vm-" + appid + ".json")
			log.Println(output)
		}

	}
}

func ajust(namespace string) {
	if len(versions) == 0 {
		var tmp1 string
		var tmp2 string
		fmt.Println("Please enter canary app's current version and new version")
		fmt.Scanln(&tmp1, &tmp2)
		versions = append(versions, tmp1)
		versions = append(versions, tmp2)
	}
	var v1 int
	var v2 int
	fmt.Println("Please input percent of", versions)
	fmt.Scanln(&v1, &v2)
	if v1+v2 != 100 {
		log.Fatal("Not Full 100%")
	}
	log.Println(map[string]int{versions[0]: v1, versions[1]: v2})
	updateVirtualService(appid, map[string]int{versions[0]: v1, versions[1]: v2}, port)
	can := ""
	fmt.Println("Ajust ", appid, "'s ", version, "canary version ?")
	fmt.Scanln(&can)
	if can == "y" || can == "Y" {
		log.Println("kubectl apply -f ./tmp/hui-" + appid + ".json")
		if compressStr(namespace) != "" {
			output := run_command("kubectl apply -f ./tmp/hui-" + appid + ".json -n " + compressStr(namespace))
			log.Println(output)
		} else {
			output := run_command("kubectl apply -f ./tmp/hui-" + appid + ".json")
			log.Println(output)
		}
	}
}

func main() {
	check()
	// 发布灰度版本,默认：v1 100%
	if mode == "canary" {
		canary(namespace)
	}
	// 调节 v1,v2 流量比例
	if mode == "ajust" {
		ajust(namespace)
	}

	// 发布v1,主要测试使用
	if mode == "normal" {
		normal(namespace)
	}
	// 完整测试流程
	if mode == "test" {
		normal(namespace)
		canary(namespace)
		ajust(namespace)
	}
}
