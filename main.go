package main

import (
	"github.com/ovirt/go-ovirt/examples"
	)

func main() {
	examples.ListDatacenters()
	examples.ListClusters()
	//examples.AddDatacenter()
	//examples.ListDatacenters()
	examples.AddCluster()
	examples.ListClusters()
}
