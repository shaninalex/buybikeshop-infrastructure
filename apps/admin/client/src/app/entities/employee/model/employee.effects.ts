import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import {
    actionEmployeeCreate,
    actionEmployeeCreateComplete, actionEmployeeCreateError,
    actionEmployeeGetList,
    actionEmployeeSetList, actionEmployeeUpdate, actionEmployeeUpdateComplete, actionEmployeeUpdateError,
} from './employee.actions';
import { catchError, exhaustMap, map, of, switchMap } from 'rxjs';
import { EmployeeApi } from '../api/api.service';
import { ApiError } from '@shared/models';
import { HttpErrorResponse } from '@angular/common/http';

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
                    catchError((errors: HttpErrorResponse) => of(actionEmployeeCreateError({errors: errors.error.errors ?? []}))),
                ),
            ),
        ),
    );

    update_employees$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionEmployeeUpdate),
            exhaustMap(action =>
                this.employeesApi.UpdateEmployee(action.id, action.data).pipe(
                    map(employee => actionEmployeeUpdateComplete({ employee })),
                    catchError((errors: HttpErrorResponse) => of(actionEmployeeUpdateError({errors: errors.error.errors ?? []}))),
                ),
            ),
        ),
    );
}
