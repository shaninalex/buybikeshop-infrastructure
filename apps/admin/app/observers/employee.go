package observers

import (
	"buybikeshop/apps/admin/app/services/employee"
	empl "buybikeshop/gen/grpc-buybikeshop-go/employee"
	"buybikeshop/libs/go/bus"
	"buybikeshop/libs/go/connector"
	"context"
	"errors"
)

var (
	ErrorSaveEmployeeDepartment  = errors.New("unable to save employee department")
	ErrorSaveEmployeeInvalidData = errors.New("invalid employee data")
)

type EmployeeObserver struct {
	bus        bus.Bus
	datasource *connector.DatasourceClient
}

func ProvideEmployeeObserver(bus bus.Bus, datasource *connector.DatasourceClient) *EmployeeObserver {
	s := &EmployeeObserver{
		bus:        bus,
		datasource: datasource,
	}
	s.init()

	return s
}

func (s *EmployeeObserver) init() {
	s.bus.Subscribe(bus.EmployeeCreatedEventType, s.employeeCreatedHandler)
}

func (s *EmployeeObserver) employeeCreatedHandler(ctx context.Context, data any) error {
	dataT, ok := data.(employee.CreateEmployeeAfter)
	if !ok {
		return ErrorSaveEmployeeInvalidData
	}

	resp, err := s.datasource.EmployeeClient.SaveEmployee(ctx, &empl.SaveEmployeeRequest{
		EmployeeId: dataT.Identity.GetId(),
		Department: dataT.Department,
	})
	if err != nil {
		return err
	}
	if resp.Employee == nil {
		return ErrorSaveEmployeeDepartment
	}

	return nil
}
