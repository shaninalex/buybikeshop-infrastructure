import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { FormInputError, UiService } from '@shared/ui';
import { filter, map, Observable, tap } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { EmployeeCreateFormModel, EmployeeViewModel } from '@entities/employee/model/employee.model';
import { form, FormField } from '@angular/forms/signals';
import { actionEmployeeUpdate, actionEmployeeUpdateComplete } from '@entities/employee/model/employee.actions';
import { Store } from '@ngrx/store';
import { employeeFormValidation } from '@entities/employee';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { provideNativeDateAdapter } from '@angular/material/core';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { Actions, ofType } from '@ngrx/effects';

@Component({
    selector: 'app-employee-detail-page',
    imports: [
        AsyncPipe,
        RouterLink,
        FormsModule,
        ReactiveFormsModule,
        FormField,
        FormInputError,
        MatFormFieldModule, MatInputModule, MatDatepickerModule, MatButtonModule
    ],
    providers: [provideNativeDateAdapter()],
    changeDetection: ChangeDetectionStrategy.OnPush,
    templateUrl: './detail-page.html',
})
export class DetailPage implements OnInit {
    private actions$ = inject(Actions);
    private route = inject(ActivatedRoute);
    private store = inject(Store);
    private ui = inject(UiService);
    private id: string;
    loading = false;

    employee$: Observable<EmployeeViewModel> = this.route.data.pipe(
        filter((data) => !!data['employee']),
        map((data) => data['employee'] as EmployeeViewModel),
        tap((employee) => {
            this.ui.setPageTitle(`Employee: ${employee.name}`);
            this.id = employee.id;
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

    ngOnInit() {
        this.actions$.pipe(
            ofType(actionEmployeeUpdateComplete),
            tap((action) => {
                this.employeeForm().reset()
                this.employeeFormModel.set({
                    name: action.employee.identity.traits.name,
                    email: action.employee.identity.traits.email,
                    phone: action.employee.identity.traits.phone,
                    dob: action.employee.identity.traits.dob,
                    photo: action.employee.identity.traits.photo,
                })
            })
        ).subscribe()
    }

    employeeForm = form(this.employeeFormModel, (schemaPath) => employeeFormValidation(schemaPath));

    submit(event: Event): void {
        event.preventDefault();
        this.loading = true;
        this.store.dispatch(
            actionEmployeeUpdate({id: this.id, data: this.employeeFormModel()}),
        );
    }
}
