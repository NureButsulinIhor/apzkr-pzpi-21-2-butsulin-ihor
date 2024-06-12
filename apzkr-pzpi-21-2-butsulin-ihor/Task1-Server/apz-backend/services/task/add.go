package task

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
	"time"
)

type BySlotGetter interface {
	GetTasksByFromSlotID(slotID uint) ([]models.Task, error)
	GetTasksByToSlotID(slotID uint) ([]models.Task, error)
}

type WarehouseGetter interface {
	GetWarehouse(id uint) (*models.Warehouse, error)
}

type SlotGetter interface {
	GetSlot(id uint) (*models.Slot, error)
}

type Adder interface {
	AddTask(task models.Task) error
}

type TransfersByWarehouseGetter interface {
	GetTransfersByWarehouseID(warehouseID uint) ([]models.Transfer, error)
}

func Add(fromSlotID uint, toSlotID uint, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.task.Add"),
	)

	l.Debug("processing auth data")
	_, payload, err := jwtauth.FromContext(cfg.Context)
	if err != nil {
		l.Debug("err to parse jwt")
		return errors.New("invalid JWT token")
	}

	user, err := models.NewUserFromClaims(payload)
	if err != nil {
		l.Debug("err to parse claims")
		return errors.New("invalid claims")
	}

	if user.Type != models.ManagerType {
		l.Debug("user is not manager")
		return errors.New("user is not manager")
	}

	l.Debug("getting manager from db")
	manager, err := cfg.Storage.GetManagerByUserID(user.ID)
	if err != nil {
		l.Debug("err to get manager from db")
		return errors.New("no manager found")
	}

	l.Debug("getting warehouse from db")
	warehouse, err := cfg.Storage.GetWarehouse(manager.WarehouseID)
	if err != nil {
		l.Error("err to get warehouse from db")
		return errors.New("internal error")
	}

	l.Debug("getting slots from db")
	fromSlot, err := cfg.Storage.GetSlot(fromSlotID)
	if err != nil {
		l.Error("err to get slots from db")
		return errors.New("no fromSlot found")
	}

	toSlot, err := cfg.Storage.GetSlot(toSlotID)
	if err != nil {
		l.Error("err to get slots from db")
		return errors.New("no toSlot found")
	}

	if fromSlot.Item == nil || toSlot.Item != nil {
		l.Debug("slots are not empty")
		return errors.New("slots are not empty")
	}

	if fromSlot.Item.Weight > toSlot.MaxWeight {
		l.Debug("weight of item is greater than max weight")
		return errors.New("weight of item is greater than max weight")
	}

	if !checkSlot(*fromSlot, *warehouse, cfg) || !checkSlot(*toSlot, *warehouse, cfg) {
		l.Debug("slots are not in warehouse")
		return errors.New("slots are not in warehouse")
	}

	l.Debug("choosing worker for task")
	worker, err := chooseWorker(warehouse.Workers, cfg)
	if err != nil {
		l.Error("err to choose worker", slog.String("error", err.Error()))
		return errors.New("no worker found")
	}

	task := models.Task{
		FromSlotID: fromSlotID,
		ToSlotID:   toSlotID,
		WorkerID:   worker.ID,
	}
	err = cfg.Storage.AddTask(task)
	if err != nil {
		l.Error("err to add task", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	return nil
}

func checkSlot(slot models.Slot, warehouse models.Warehouse, cfg Configuration) bool {
	l := cfg.Logger.With(
		slog.String("op", "services.task.checkSlot"),
	)

	l.Debug("checking slot for availability")
	if slot.StorageID == warehouse.StorageID {
		return true
	}

	transfers, err := cfg.Storage.GetTransfersByWarehouseID(warehouse.ID)
	if err != nil {
		l.Error("err to get transfers from db", slog.String("error", err.Error()))
		return false
	}

	for _, transfer := range transfers {
		if transfer.InDate.Before(time.Now()) &&
			transfer.OutDate.After(time.Now()) &&
			transfer.Warehouse.StorageID == slot.StorageID {
			return true
		}
	}

	l.Debug("checking slot for not being busy")
	tasks, err := cfg.Storage.GetTasksByFromSlotID(slot.ID)
	if err != nil {
		l.Error("err to get tasks from db", slog.String("error", err.Error()))
	}
	outTasks, err := cfg.Storage.GetTasksByToSlotID(slot.ID)
	if err != nil {
		l.Error("err to get tasks from db", slog.String("error", err.Error()))
	}
	tasks = append(tasks, outTasks...)

	return len(tasks) == 0
}

func chooseWorker(workers []models.Worker, cfg Configuration) (models.Worker, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.task.chooseWorker"),
	)

	var workingWorkers []models.Worker
	for _, worker := range workers {
		if worker.IsWorking() {
			workingWorkers = append(workingWorkers, worker)
		}
	}

	var chosenWorker models.Worker
	countTasks := uint(0)
	countTasks--
	for _, worker := range workingWorkers {
		tasks, err := cfg.Storage.GetTasksByWorkerID(worker.ID)
		if err != nil {
			l.Error("err to get tasks from db", slog.String("error", err.Error()))
			continue
		}
		if countTasks > uint(len(tasks)) {
			chosenWorker = worker
			countTasks = uint(len(tasks))
		}
	}

	if chosenWorker.ID == 0 {
		return chosenWorker, errors.New("no worker found")
	}

	return chosenWorker, nil
}
