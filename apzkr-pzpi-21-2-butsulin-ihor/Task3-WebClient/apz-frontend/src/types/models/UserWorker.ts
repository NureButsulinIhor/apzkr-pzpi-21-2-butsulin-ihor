type UserWorker = {
    ID: bigint;
    userID: bigint;
    user: User;
    warehouseID: bigint;
    timetables: Timetable[];
}