import { createReducer, on } from '@ngrx/store';
import { EmployeeModel } from './employee.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { actionEmployeeSetList } from './employee.actions';

export interface EmployeeState extends EntityState<EmployeeModel> {}

export const employeesAdapter = createEntityAdapter<EmployeeModel>({
    selectId: (employee: EmployeeModel) => employee.identity.id,
});
export const initialState = employeesAdapter.getInitialState();

export const employeeReducer = createReducer(
    initialState,
    on(actionEmployeeSetList, (state, action) => employeesAdapter.addMany(action.employees, state)),
);
