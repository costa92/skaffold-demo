package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		}

		svc, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("There are %d svc in the cluster\n", len(svc.Items))
		}
		time.Sleep(10 * time.Second)
	}
}

func printJson(obj interface{}, err error) {
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(obj)
	if err != nil {
		log.Fatal("序列化失败!", err)
	}
  log.Println("printJson")
	fmt.Println(string(data))
}
