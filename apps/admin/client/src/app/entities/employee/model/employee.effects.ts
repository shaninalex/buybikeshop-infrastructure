import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { actionEmployeeGetList, actionEmployeeSetList, } from './employee.actions';
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
}
