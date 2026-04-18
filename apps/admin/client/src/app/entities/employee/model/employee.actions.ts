import { createAction, props } from '@ngrx/store';
import { EmployeeCreateFormModel, EmployeeModel } from './employee.model';

export const actionEmployeeGetList = createAction('[Employee] get list');
export const actionEmployeeSetList = createAction(
    '[Employee] set list',
    props<{ employees: EmployeeModel[] }>(),
);


export const actionEmployeeCreate = createAction(
    '[Employee] create',
    props<{ data: EmployeeCreateFormModel }>(),
);

export const actionEmployeeCreateError = createAction(
    '[Employee] create error',
    props<{ errors: any }>(),
);

export const actionEmployeeCreateComplete = createAction(
    '[Employee] create complete',
    props<{ employee: EmployeeModel }>(),
);
