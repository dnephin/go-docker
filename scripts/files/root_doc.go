/*
Package docker is the official Go client for the Docker API.

The "docker" command uses this package to communicate with the daemon. It can also
be used by your own Go applications to do anything the command-line interface does
- running containers, pulling images, managing swarms, etc.

For more information about the Engine API, see the documentation:
https://docs.docker.com/engine/reference/api/

Usage

You use the library by creating a client object and calling methods on it. The
client can be created either from environment variables with NewEnvClient, or
configured manually with NewClient.

For example, to list running containers (the equivalent of "docker ps"):

	package main

	import (
		"context"
		"fmt"

		"golang.docker.io/go-docker/api/types"
		"golang.docker.io/go-docker"
	)

	func main() {
		cli, err := docker.NewEnvClient()
		if err != nil {
			panic(err)
		}

		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

		for _, container := range containers {
			fmt.Printf("%s %s\n", container.ID[:10], container.Image)
		}
	}

Dependency management

This package uses the https://github.com/golang/dep tool, and dependencies
are defined in Gopkg.toml and Gopkg.lock files. After writing a piece of code
importing this package, you can use the following to vendor for the first time
the correct dependencies:

	dep init

In order to update the dependency, modify the Gopkg.toml to point to the desired
branch or tag and run:

	dep ensure

*/
package docker // import "golang.docker.io/go-docker"
