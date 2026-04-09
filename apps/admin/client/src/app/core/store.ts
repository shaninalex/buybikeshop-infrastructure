import { EmployeeEffects, employeeReducer } from '@entities/employee';

export const effects = [EmployeeEffects];

export const reducers = {
    employee: employeeReducer,
};
