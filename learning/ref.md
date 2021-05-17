
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
  - https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/autoscaling#Client
  - https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/ec2@v1.6.0/types#NetworkInterface





- aws-skd-go
  - https://gist.github.com/eferro/651fbb72851fa7987fc642c8f39638eb
  - https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#DescribeInstancesInput
    - network-interface.private-dns-name

- go with json
  - https://yourbasic.org/golang/json-example/
  - https://www.sohamkamani.com/golang/parsing-json/
