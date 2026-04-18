import { createFeatureSelector, createSelector } from '@ngrx/store';
import { employeesAdapter, EmployeeState } from './employee.reducer';
import { IdentityTraits } from './employee.model';

export const selectEmployeesFeature = createFeatureSelector<EmployeeState>('employee');
export const employeesSelectors = employeesAdapter.getSelectors();

export const selectEmployees = createSelector(selectEmployeesFeature, (state) =>
    employeesSelectors.selectAll(state),
);

export const selectEmployeeById = (id: string) =>
    createSelector(selectEmployeesFeature, (state: EmployeeState) =>
        employeesSelectors.selectAll(state).find((p) => p.identity.id === id),
    );

export const selectEmployeeViewModel = (id: string) =>
    createSelector(selectEmployeeById(id), employee => {
        if (!employee) return null;

        const traits = employee.identity.traits as IdentityTraits;
        const created_at = employee.identity.created_at? new Date(employee.identity.created_at) : null;

        return {
            id: employee.identity.id,
            email: traits.email,
            name: traits.name,
            photo: traits.photo,
            dob: traits.dob,
            phone: traits.phone,
            state: employee.identity.state,
            created_at: created_at,
        };
    });
