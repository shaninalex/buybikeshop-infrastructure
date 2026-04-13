import { email, form, FormField, pattern, required } from '@angular/forms/signals';
import { Component, DestroyRef, inject, OnInit, signal } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { EmployeeCreateFormModel } from '@entities/employee/model/employee.model';
import { Store } from '@ngrx/store';
import {
    actionEmployeeCreate,
    actionEmployeeCreateComplete,
    actionEmployeeCreateError
} from '@entities/employee/model/employee.actions';
import { Actions, ofType } from '@ngrx/effects';
import { tap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { FormInputError, PasswordGenerator } from '@shared/ui';

@Component({
    selector: 'app-employee-create-manual-form',
    imports: [
        FormsModule,
        ReactiveFormsModule,
        FormField,
        FormInputError,
        PasswordGenerator
    ],
    template: `
        <form (submit)="submit($event)">
            <div class="d-flex justify-content-between align-items-center mb-4">
                <h5 class="mb-3">Employee Information</h5>
                <div class="btn-group">
                    <button type="button" class="btn btn-outline-secondary btn-sm" (click)="employeeForm().reset()">
                        Cancel
                    </button>
                    <button type="submit" class="btn btn-primary btn-sm"
                            [disabled]="employeeForm().invalid() || loading">
                        @if (loading) {
                            <i class="fa-solid fa-circle-notch fa-spin"></i>
                        }
                        Create Employee
                    </button>
                </div>
            </div>

            <div class="row g-3">
                <div class="col-md-6">
                    <label class="form-label" for="name">Full Name *</label>
                    <input
                        type="text"
                        id="name"
                        class="form-control"
                        [formField]="employeeForm.name"
                        placeholder="John Doe"
                    />
                    <app-form-input-error [inputField]="employeeForm.name" />
                </div>

                <div class="col-md-6">
                    <label class="form-label" for="email">Email *</label>
                    <input
                        type="email"
                        id="email"
                        class="form-control"
                        [formField]="employeeForm.email"
                        placeholder="john.doe@company.com"
                    />
                    <app-form-input-error [inputField]="employeeForm.email" />
                </div>

                <div class="col-md-6">
                    <label class="form-label" for="phone">Phone</label>
                    <input
                        type="tel"
                        id="phone"
                        class="form-control"
                        [formField]="employeeForm.phone"
                    />
                    <app-form-input-error [inputField]="employeeForm.phone" />
                </div>

                <div class="col-md-6">
                    <label class="form-label" for="dob">Date of Birth</label>
                    <input
                        type="text"
                        id="dob"
                        class="form-control"
                        placeholder="YYYY-MM-DD"
                        [formField]="employeeForm.dob"
                    />
                    <app-form-input-error [inputField]="employeeForm.dob" />
                </div>

                <div class="col-12">
                    <label class="form-label" for="photo">Photo URL</label>
                    <input
                        type="url"
                        id="photo"
                        class="form-control"
                        [formField]="employeeForm.photo"
                        placeholder="https://example.com/photo.jpg"
                    />
                    <app-form-input-error [inputField]="employeeForm.photo" />
                </div>
            </div>

            <hr class="my-4"/>

            <!-- Password -->
            <h5 class="mb-3">Security</h5>
            <div class="row g-3">
                <div class="col-md-6">
                    <label class="form-label" for="password">Password *</label>
                    <div class="d-flex align-items-center justify-content-start gap-2">
                        <input
                            [type]="passwordType ? 'password' : 'text'"
                            id="password"
                            class="form-control"
                            [formField]="employeeForm.password"
                            placeholder="Enter password"
                        />
                        <button class="btn btn-sm btn-outline-secondary" type="button" (click)="passwordTypeToggle()">
                            @if (passwordType) {
                                <i class="fa-regular fa-eye"></i>
                            } @else {
                                <i class="fa-solid fa-eye-slash"></i>
                            }
                        </button>
                        <app-password-generator (password)="passwordGenerated($event)" />
                    </div>
                    <app-form-input-error [inputField]="employeeForm.password" />
                </div>
            </div>
        </form>
    `,
})
export class ManualForm implements OnInit {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    loading = false;
    passwordType = true;

    ngOnInit() {
        this.actions$.pipe(
            ofType(actionEmployeeCreateComplete), takeUntilDestroyed(this.destroyRef)
        ).subscribe(() => (this.loading = false));
        this.actions$.pipe(
            ofType(actionEmployeeCreateError),
            tap(action => {
                console.log(action.errors)
            }),
            takeUntilDestroyed(this.destroyRef),
        ).subscribe(() => (this.loading = false));
    }

    employeeFormModel = signal<EmployeeCreateFormModel>({
        name: '',
        email: '',
        phone: '',
        dob: '',
        photo: '',
        password: ''
    })

    employeeForm = form(this.employeeFormModel, (schemaPath) => {
        email(schemaPath.email, {message: 'Invalid email address format'});

        required(schemaPath.email, {message: 'Email is required'});
        required(schemaPath.password, {message: 'Password is required'});

        required(schemaPath.name, {message: 'Name is required'});
        pattern(
            schemaPath.name,
            /^[\p{L} ]+$/u,
            { message: 'Name may contain only letters and spaces' }
        );

        pattern(
            schemaPath.dob,
            /^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$/,
            { message: 'Date of birth must be in YYYY-MM-DD format' }
        );

        // Phone validator: digits, +, (), space, and dash
        pattern(
            schemaPath.phone,
            /^$|^[0-9+()\- ]+$/,
            {
                message: 'Phone number may contain only digits, spaces, "+", "-", and parentheses'
            }
        );
    });

    submit(event: Event): void {
        event.preventDefault();
        this.loading = true;
        this.store.dispatch(
            actionEmployeeCreate({data: this.employeeFormModel()}),
        );
    }

    passwordGenerated(pswd: string) {
        this.employeeForm.password().value.set(pswd);
    }

    passwordTypeToggle() {
        this.passwordType = !this.passwordType;
    }
}
