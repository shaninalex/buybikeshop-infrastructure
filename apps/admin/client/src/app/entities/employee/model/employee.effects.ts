import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import {
    actionEmployeeCreate,
    actionEmployeeCreateComplete, actionEmployeeCreateError,
    actionEmployeeGetList,
    actionEmployeeSetList,
} from './employee.actions';
import { catchError, exhaustMap, map, of, switchMap } from 'rxjs';
import { EmployeeApi } from '../api/api.service';
import { ApiError } from '@shared/models';

@Injectable()
export class EmployeeEffects {
    private actions$ = inject(Actions);
    private employeesApi = inject(EmployeeApi);

    get_employees_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionEmployeeGetList),
            exhaustMap(() =>
                this.employeesApi.GetEmployees().pipe(
                    map((data) => actionEmployeeSetList({ employees: data })),
                    // catchError((errors: ApiError[]) => of(actionEmployeeSetListError({ errors: errors ?? [] }))),
                ),
            ),
        ),
    );

    create_employees$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionEmployeeCreate),
            exhaustMap(action =>
                this.employeesApi.CreateEmployee(action.data).pipe(
                    map(employee => actionEmployeeCreateComplete({ employee })),
                    catchError((errors: ApiError[]) => of(actionEmployeeCreateError({ errors: errors ?? [] }))),
                ),
            ),
        ),
    );
}
