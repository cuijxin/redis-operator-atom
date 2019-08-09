module github.com/cuijxin/redis-operator-atom

go 1.12

require (
	github.com/cuijxin/kooper-atom v0.0.0-20190808091322-f41fd41f7eff
	github.com/grpc-ecosystem/grpc-gateway v1.6.2 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.3.0
	k8s.io/api v0.0.0-20190802060718-d0d4f3afa3ab
	k8s.io/apiextensions-apiserver v0.0.0-20190802061903-25691aabac0a
	k8s.io/apimachinery v0.0.0-20190802060556-6fa4771c83b3
	k8s.io/client-go v11.0.0+incompatible
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible
)
