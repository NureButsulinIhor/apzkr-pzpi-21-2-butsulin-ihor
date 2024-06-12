type Warehouse = {
    ID: bigint
    storageID: bigint
    storage: ItemStorage
    workers: UserWorker[]
    manager: Manager
}