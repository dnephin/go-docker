// +build linux freebsd solaris openbsd darwin

package docker // import "golang.docker.com/go-docker"

// DefaultDockerHost defines os specific default if DOCKER_HOST is unset
const DefaultDockerHost = "unix:///var/run/docker.sock"
