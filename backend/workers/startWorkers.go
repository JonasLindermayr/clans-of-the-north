package workers

func StartWorkers() {
	go ResourceProdWorker()
	go SaveCacheToDB()
}