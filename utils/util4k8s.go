package utils

import (
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"math"
	"strconv"

	"os"
	"path/filepath"
)


func CreateK8SClient() (client *k8s.Clientset, err error){

	h := homeDir()
	h = filepath.Join(h, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", h)
	if err != nil{
		panic(err)
	}

	client, err = k8s.NewForConfig(config)
	if err != nil {
		 panic(err)
	}
	return
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != ""{
		return h
	}
	return os.Getenv("USERPROFILE") // for windows
}

//取整和并将float64 转为 string
func FltoStr(f float64) string{

	f = math.Ceil(f)

	t := strconv.FormatFloat(f,'f',-1,64)
	return t
}