type Task = {
    ID: bigint;
    workerID: bigint;
    worker: UserWorker;
    fromSlotID: bigint;
    fromSlot: Slot;
    toSlotID: bigint;
    toSlot: Slot;
    status: boolean
}