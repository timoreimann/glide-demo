package main

import "k8s.io/kubernetes/pkg/client/restclient"

func main() {
	_ = &restclient.Config{}
}
