module github.com/polynetwork/fabric-relayer

go 1.14

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/boltdb/bolt v1.3.1
	github.com/cloudflare/cfssl v1.4.1
	github.com/cmars/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-kit/kit v0.10.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.5.0
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/hyperledger/fabric-config v0.0.5
	github.com/hyperledger/fabric-lib-go v1.0.0
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta3.0.20201006151309-9c426dcc5096
	github.com/kr/pretty v0.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/miekg/pkcs11 v1.0.3
	github.com/mitchellh/mapstructure v1.3.2
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/gomega v1.10.3 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/polynetwork/poly v0.0.0-20200722075529-eea88acb37b2
	github.com/polynetwork/poly-go-sdk v0.0.0-20200729103825-af447ef53ef0
	github.com/prometheus/client_golang v1.7.1
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.6.1
	github.com/tjfoc/gmsm v1.3.2-0.20200914155643-24d14c7bd05c
	github.com/urfave/cli v1.22.4
	go.etcd.io/bbolt v1.3.5 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/genproto v0.0.0-20201110150050-8816d57aaa9a // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)

replace (
	github.com/go-kit/kit v0.10.0 => github.com/go-kit/kit v0.8.0
	github.com/polynetwork/poly => github.com/zouxyan/poly v0.0.0-20201110080649-bde9b073a9fc
	github.com/tjfoc/gmsm => github.com/zouxyan/gmsm v1.3.2-0.20200925082225-a66aabdb8da8
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
)
