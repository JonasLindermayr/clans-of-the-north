package workers

func StartWorkers() {
	go ResourceProdWorker()
	go BuildingQueueWorker()
	go SaveCacheToDB()
}