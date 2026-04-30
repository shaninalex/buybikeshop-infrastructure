package observers

import (
	"buybikeshop/apps/admin/app/services/employee"
	"buybikeshop/libs/go/bus"
	"buybikeshop/libs/go/keto"
	"context"
	"errors"
)

type PermissionObserver struct {
	bus               bus.Bus
	permissionService *keto.Manager
}

func ProvidePermissionObserver(bus bus.Bus, permissionService *keto.Manager) *PermissionObserver {
	s := &PermissionObserver{
		bus:               bus,
		permissionService: permissionService,
	}
	s.init()

	return s
}

func (s *PermissionObserver) init() {
	s.bus.Subscribe(bus.EmployeeCreatedEventType, s.employeeCreatedHandler)
}

func (s *PermissionObserver) employeeCreatedHandler(ctx context.Context, data any) error {
	dataT, ok := data.(employee.CreateEmployeeAfter)
	if !ok {
		return errors.New("invalid employee permission data")
	}

	id := dataT.Identity.GetId()
	if err := s.permissionService.Assign(ctx, dataT.Group, &id, nil); err != nil {
		return err
	}

	return nil
}
