import { Component, inject, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { UiService } from '@shared/ui';
import { filter, map, Observable, tap } from 'rxjs';
import { AsyncPipe, DatePipe } from '@angular/common';
import { EmployeeCreateFormModel, EmployeeViewModel } from '@entities/employee/model/employee.model';
import { form } from '@angular/forms/signals';
import { actionEmployeeCreate } from '@entities/employee/model/employee.actions';
import { Store } from '@ngrx/store';
import { employeeFormValidation } from '@entities/employee';

@Component({
    selector: 'app-employee-detail-page',
    imports: [
        AsyncPipe,
        RouterLink,
        DatePipe,
    ],
    templateUrl: './detail-page.html',
})
export class DetailPage {
    private route = inject(ActivatedRoute);
    private store = inject(Store);
    private ui = inject(UiService);
    loading = false;
    tmpState: EmployeeViewModel;

    employee$: Observable<EmployeeViewModel> = this.route.data.pipe(
        filter((data) => !!data['employee']),
        map((data) => data['employee'] as EmployeeViewModel),
        tap((employee) => {
            this.tmpState = {...employee};
            this.ui.setPageTitle(`Employee: ${employee.name}`);
            this.employeeFormModel.set({
                name: employee.name,
                email: employee.email,
                phone: employee.phone,
                dob: employee.dob,
                photo: employee.photo,
            })
        }),
    );

    employeeFormModel = signal<EmployeeCreateFormModel>({
        name: '',
        email: '',
        phone: '',
        dob: new Date(),
        photo: '',
    })

    employeeForm = form(this.employeeFormModel, (schemaPath) => employeeFormValidation(schemaPath));

    submit(event: Event): void {
        event.preventDefault();
        this.loading = true;
        this.store.dispatch(
            actionEmployeeCreate({data: this.employeeFormModel()}),
        );
    }
}
