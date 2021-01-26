
func getNodeEvents(c kubernetes.Interface, nodeName string) []v1.Event {
	selector := fields.Set{
		"involvedObject.kind":      "Node",
		"involvedObject.name":      nodeName,
		"involvedObject.namespace": metav1.NamespaceAll,
		"source":                   "kubelet",
	}.AsSelector().String()
	options := metav1.ListOptions{FieldSelector: selector}
	events, err := c.CoreV1().Events(metav1.NamespaceSystem).List(options)
	if err != nil {
		Logf("Unexpected error retrieving node events %v", err)
		return []v1.Event{}
	}
	return events.Items
}

// nodes, err := client.CoreV1().Nodes().List(ontext.TODO(), metav1.ListOptions{FieldSelector: "metadata.name=cluster"})
// nodes, err := client.CoreV1().Nodes().List(ontext.TODO(), metav1.ListOptions{LabelSelector: "app=<APPNAME>"})


for _, deploy := range deployments.Items {
    fmt.Println(deploy.Name, deploy.CreationTimestamp)
    fmt.Println(deploy.ObjectMeta.GetCreationTimestamp(), deploy.ObjectMeta.GetLabels())

    //      printDeploymentSpecYaml(deploy)

}
