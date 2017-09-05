//go:generate protoc -I . --gogofast_out=import_path=golang.docker.com/go-docker/api/types/swarm/runtime:. plugin.proto

package runtime
