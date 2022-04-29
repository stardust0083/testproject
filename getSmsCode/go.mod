module getSmsCode

go 1.16

require (
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20220407094043-a94812496cf5 // indirect
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20220428095506-8293988499ab // indirect
	github.com/asim/go-micro/plugins/transport/grpc/v4 v4.0.0-20220428095506-8293988499ab // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/miekg/dns v1.1.48 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/urfave/cli/v2 v2.5.0 // indirect
	github.com/xanzy/ssh-agent v0.3.1 // indirect
	go-micro.dev/v4 v4.6.0
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	google.golang.org/protobuf v1.28.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
