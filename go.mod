module github.com/ocibuilder/lib

go 1.13

require (
	cloud.google.com/go v0.38.0
	github.com/Azure/azure-storage-blob-go v0.8.0
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v2.0.7+incompatible
	github.com/artbegolli/yenv v0.2.0
	github.com/aws/aws-sdk-go v1.30.4
	github.com/containerd/containerd v1.3.0 // indirect
	github.com/docker/docker v1.14.0-0.20190319215453-e7b5f7dbe98c
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/gobuffalo/packr v1.30.1
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/uuid v1.1.1
	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/k14s/ytt v0.26.0
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/moby/buildkit v0.6.3
	github.com/nwaples/rardecode v1.1.0 // indirect
	github.com/ocibuilder/gofeas v1.0.0
	github.com/ocibuilder/ocibuilder v0.0.0-20200401153708-0390955eae15
	github.com/pierrec/lz4 v2.4.1+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.5.1
	github.com/tidwall/gjson v1.6.0
	github.com/tidwall/sjson v1.1.1
	github.com/ulikunitz/xz v0.5.7 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	go.opencensus.io v0.22.3 // indirect
	golang.org/x/crypto v0.0.0-20191112222119-e1110fd1c708
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	google.golang.org/api v0.21.0
	gopkg.in/src-d/go-git.v4 v4.13.1
	k8s.io/api v0.18.0
	k8s.io/apimachinery v0.18.0
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8 // indirect
)

replace github.com/containerd/containerd v1.3.0-0.20190507210959-7c1e88399ec0 => github.com/containerd/containerd v1.3.0

replace github.com/docker/docker v1.14.0-0.20190319215453-e7b5f7dbe98c => github.com/docker/docker v1.4.2-0.20191113233703-44d39013868a

replace golang.org/x/crypto v0.0.0-20190129210102-0709b304e793 => golang.org/x/crypto v0.0.0-20191112222119-e1110fd1c708
