// 项目的总入口
package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "km_backend/config"
	"km_backend/utils/logs"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	option         string = "delete"    // option: add, delete, update, get
	resource       string = "namespace" // resource: namespace, deployment
	resourceName   string = "test001"   // resourceName: the name of the namespace or deployment to be added or deleted
	resourceNsName string = "test001"   // resourceNsName: the name of the namespace where the deployment is located, only used when resource is deployment

	newNamespace corev1.Namespace  // 定义一个新的namespace对象
	newDeopyment appsv1.Deployment // 定义一个新的deployment对象
)

// 增加一个namespace
func addNamespace(clientset *kubernetes.Clientset, resourceName string) {
	newNamespace.Name = resourceName
	_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &newNamespace, metav1.CreateOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"Error:": err.Error()}, "namespace创建失败.")
	}
}

// 增加一个deployment
func addDeployment(clientset *kubernetes.Clientset, resourceName string) {
	newDeopyment.Name = resourceName
	newDeopyment.Spec.Replicas = new(int32)
	*newDeopyment.Spec.Replicas = 3
	newDeopyment.Spec.Template.Spec.Containers = append(newDeopyment.Spec.Template.Spec.Containers, corev1.Container{
		Name:  "nginx",
		Image: "nginx:latest",
	})

	label := make(map[string]string)
	label["app"] = "nginx"
	label["version"] = "v1"
	newDeopyment.Spec.Selector = &metav1.LabelSelector{}
	newDeopyment.Spec.Selector.MatchLabels = label
	newDeopyment.Spec.Template.ObjectMeta.Labels = label

	_, err := clientset.AppsV1().Deployments(resourceNsName).Create(context.TODO(), &newDeopyment, metav1.CreateOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"Error:": err.Error()}, "deployment创建失败.")
	}
}

func addDeploymentForJson(clientset *kubernetes.Clientset, resourceName string) {
	deployJson := `{
		"kind": "Deployment",
		"apiVersion": "apps/v1",
		"metadata": {
			"name": "redis",
			"creationTimestamp": null,
			"labels": {
				"app": "redis"
			}
		},
		"spec": {
			"replicas": 1,
			"selector": {
				"matchLabels": {
					"app": "redis"
				}
			},
			"template": {
				"metadata": {
					"creationTimestamp": null,
					"labels": {
						"app": "redis"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "redis",
							"image": "redis",
							"resources": {}
						}
					]
				}
			},
			"strategy": {}
		},
		"status": {}
	}`
	err := json.Unmarshal([]byte(deployJson), &newDeopyment)
	if err != nil {
		logs.Error(map[string]interface{}{"Error:": err.Error()}, "deployment json解析失败.")
	}
	_, err = clientset.AppsV1().Deployments(resourceNsName).Create(context.TODO(), &newDeopyment, metav1.CreateOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"Error:": err.Error()}, "deployment创建失败.")
	}
}

// 删除一个namespace
func deleteNamespace(clientset *kubernetes.Clientset, resourceName string) {
	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), resourceName, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"Error:": err.Error()}, "namespace删除失败.")
	}
}

// 删除一个deployment
func deleteDeployment(clientset *kubernetes.Clientset, resourceName string) {
	err := clientset.AppsV1().Deployments(resourceNsName).Delete(context.TODO(), resourceName, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"Error": err.Error()}, "deployment删除失败.")
	}
}

// 流程控制语句
func selectActions() {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "./config/kubeconfig")
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// select menu
	switch option {
	case "add":
		switch resource {
		case "namespace":
			addNamespace(clientset, resourceName)
		case "deployment":
			// addDeployment(clientset, resourceName)
			addDeploymentForJson(clientset, resourceName)
		}
	case "delete":
		switch resource {
		case "namespace":
			deleteNamespace(clientset, resourceName)
		case "deployment":
			deleteDeployment(clientset, resourceName)
		}
	case "update":
		fmt.Println("update")
	case "get":
		fmt.Println("get")
	default:
		fmt.Println("invalid option")
	}
}

func main() {
	selectActions()
}
