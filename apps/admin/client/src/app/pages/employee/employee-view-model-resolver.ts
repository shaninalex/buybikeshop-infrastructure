import { ActivatedRouteSnapshot, ResolveFn } from '@angular/router';
import { inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionEmployeeGetList } from '@entities/employee';
import { filter, tap } from 'rxjs';
import { selectEmployeeViewModel } from '@entities/employee/model/employee.selectors';
import { EmployeeViewModel } from '@entities/employee/model/employee.model';

export const employeeViewModelResolver: ResolveFn<EmployeeViewModel | null> = (route: ActivatedRouteSnapshot) => {
    const store = inject(Store);
    const id = route.paramMap.get('id');
    if (!id) {
        return null;
    }

    return store.select(selectEmployeeViewModel(id)).pipe(
        tap(empl => {
            if (!empl) { store.dispatch(actionEmployeeGetList()) }
        }),
        filter((empl) => !!empl)
    );
};
