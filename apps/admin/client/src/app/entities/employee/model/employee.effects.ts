import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import {
    actionEmployeeCreate,
    actionEmployeeCreateComplete,
    actionEmployeeGetList,
    actionEmployeeSetList,
} from './employee.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { EmployeeApi } from '../api/api.service';

@Injectable()
export class EmployeeEffects {
    private actions$ = inject(Actions);
    private employeesApi = inject(EmployeeApi);

    get_employees_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionEmployeeGetList),
            exhaustMap(() =>
                this.employeesApi
                    .GetEmployees()
                    .pipe(switchMap((data) => of(actionEmployeeSetList({ employees: data })))),
            ),
        ),
    );

    create_employees$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionEmployeeCreate),
            exhaustMap(action =>
                this.employeesApi
                    .CreateEmployee(action.data)
                    // TODO: handle error
                    .pipe(switchMap((data) => of(actionEmployeeCreateComplete({ employee: data })))),
            ),
        ),
    );
}
