package observers

import (
	"buybikeshop/libs/go/bus"
	"buybikeshop/libs/go/keto"
	"context"
	"errors"

	"github.com/google/uuid"
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

type EmployeeCreatedPermissionData struct {
	EmployeeID uuid.UUID
	Group      string
}

func (s *PermissionObserver) employeeCreatedHandler(ctx context.Context, data any) error {
	dataT, ok := data.(EmployeeCreatedPermissionData)
	if !ok {
		return errors.New("invalid employee permission data")
	}

	id := dataT.EmployeeID.String()
	if err := s.permissionService.Assign(ctx, dataT.Group, &id, nil); err != nil {
		return err
	}

	return nil
}
