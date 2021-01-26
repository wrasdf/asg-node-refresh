
apiType:
  Ref:
    - https://pkg.go.dev/k8s.io/api/apps/v1
    - https://pkg.go.dev/k8s.io/api/core/v1
    - https://pkg.go.dev/k8s.io/apimachinery
    - https://godoc.org/k8s.io/client-go/kubernetes#Clientset

metav1:
  - https://pkg.go.dev/github.com/ericchiang/k8s/apis/meta/v1#ObjectMeta
  - https://pkg.go.dev/github.com/ericchiang/k8s/apis/meta/v1#Time


- examples:
  - https://github.com/pachyderm/pachyderm/blob/master/src/server/pps/server/api_server.go
  - https://github.com/golang/protobuf

- aws-sdk-go-v2
  - https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/ec2#Client
