import { createAction, props } from '@ngrx/store';
import { EmployeeModel } from './employee.model';

export const actionEmployeeGetList = createAction('[Employee] get list');
export const actionEmployeeSetList = createAction(
    '[Employee] set list',
    props<{ employees: EmployeeModel[] }>(),
);
