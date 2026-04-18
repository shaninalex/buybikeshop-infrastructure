import { createAction, props } from '@ngrx/store';
import { EmployeeCreateFormModel, EmployeeModel } from './employee.model';
import { ApiError } from '@shared/models';

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
    props<{ errors: ApiError[] }>(),
);

export const actionEmployeeCreateComplete = createAction(
    '[Employee] create complete',
    props<{ employee: EmployeeModel }>(),
);

export const actionEmployeeUpdate = createAction(
    '[Employee] update',
    props<{ id: string, data: EmployeeCreateFormModel }>(),
);

export const actionEmployeeUpdateError = createAction(
    '[Employee] update error',
    props<{ errors: ApiError[] }>(),
);

export const actionEmployeeUpdateComplete = createAction(
    '[Employee] update complete',
    props<{ employee: EmployeeModel }>(),
);
