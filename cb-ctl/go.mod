module github.com/cloud-barista/cb-client/cb-ctl

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
)

require github.com/cloud-barista/cb-tumblebug v0.3.1
