package main

import "github.com/TSF/TFSMasterNameNode/internal/infrastructure/delivery/api"

func main() {
	deliveryStrategy := api.NewApi()
	deliveryStrategy.Start()
}
