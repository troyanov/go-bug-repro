package main

import (
	"context"
	"fmt"
	"os"

	"github.com/canonical/microcluster/v2/microcluster"
)

func main() {
	fmt.Println("running main")
	dir, err := os.MkdirTemp("", "go-bug-repro")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	cluster, err := microcluster.App(microcluster.Args{
		StateDir: dir,
	})

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := cluster.Start(ctx, microcluster.DaemonArgs{
			Version: "dev",
			Debug:   true,
			Verbose: true,
		})
		if err != nil {
			panic(err)
		}
	}()

	err = cluster.Ready(context.Background())
	if err != nil {
		panic(err)
	}

	err = cluster.NewCluster(ctx, "test", "127.0.0.1:5280", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("yay, cluster is bootstrapped")
}
