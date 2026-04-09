import { createFeatureSelector, createSelector } from '@ngrx/store';
import { employeesAdapter, EmployeeState } from './employee.reducer';

export const selectEmployeesFeature = createFeatureSelector<EmployeeState>('employee');
export const employeesSelectors = employeesAdapter.getSelectors();

export const selectEmployees = createSelector(selectEmployeesFeature, (state) =>
    employeesSelectors.selectAll(state),
);

export const selectEmployeeByID = (id: number) =>
    createSelector(selectEmployeesFeature, (state: EmployeeState) =>
        employeesSelectors.selectAll(state).find((p) => p.id === id),
    );
