import { email, form, FormField, required } from '@angular/forms/signals';
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

@Component({
    selector: 'app-employee-create-manual-form',
    imports: [
        FormsModule,
        ReactiveFormsModule,
        FormField
    ],
    template: `
        <form (submit)="submit($event)">
            <div class="d-flex justify-content-between align-items-center mb-4">
                <h5 class="mb-3">Employee Information</h5>
                <div class="btn-group">
                    <button type="button" class="btn btn-outline-secondary btn-sm">
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


            <ul class="error-list">
                @for (error of employeeForm().errors(); track error) {
                    <li class="text-red-500 text-sm">{{ error.message }}</li>
                }
            </ul>

            <div class="row g-3">
                <div class="col-md-6">
                    <label class="form-label">Full Name</label>
                    <input
                        type="text"
                        class="form-control"
                        [formField]="employeeForm.name"
                        placeholder="John Doe"
                    />
                    @if (employeeForm.name().touched()) {
                        <ul class="error-list">
                            @for (error of employeeForm.name().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
                </div>

                <div class="col-md-6">
                    <label class="form-label">Email</label>
                    <input
                        type="email"
                        class="form-control"
                        [formField]="employeeForm.email"
                        placeholder="john.doe@company.com"
                    />
                    @if (employeeForm.email().touched()) {
                        <ul class="error-list">
                            @for (error of employeeForm.email().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
                </div>

                <div class="col-md-6">
                    <label class="form-label">Phone</label>
                    <input
                        type="tel"
                        class="form-control"
                        [formField]="employeeForm.phone"
                    />
                    @if (employeeForm.phone().touched()) {
                        <ul class="error-list">
                            @for (error of employeeForm.phone().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
                </div>

                <div class="col-md-6">
                    <label class="form-label">Date of Birth</label>
                    <input
                        type="date"
                        class="form-control"
                        [formField]="employeeForm.dob"
                    />
                    @if (employeeForm.dob().touched()) {
                        <ul class="error-list">
                            @for (error of employeeForm.dob().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
                </div>

                <div class="col-12">
                    <label class="form-label">Photo URL</label>
                    <input
                        type="url"
                        class="form-control"
                        [formField]="employeeForm.photo"
                        placeholder="https://example.com/photo.jpg"
                    />
                    @if (employeeForm.photo().touched()) {
                        <ul class="error-list">
                            @for (error of employeeForm.photo().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
                </div>
            </div>

            <hr class="my-4"/>

            <!-- Password -->
            <h5 class="mb-3">Security</h5>
            <div class="row g-3">
                <div class="col-md-6">
                    <label class="form-label">Password</label>
                    <input
                        type="password"
                        class="form-control"
                        [formField]="employeeForm.password"
                        placeholder="Enter password"
                    />
                    @if (employeeForm.password().touched()) {
                        <ul class="error-list">
                            @for (error of employeeForm.password().errors(); track error) {
                                <li class="text-red-500 text-sm">{{ error.message }}</li>
                            }
                        </ul>
                    }
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
        required(schemaPath.email, {message: 'Email is required'});
        email(schemaPath.email, {message: 'Invalid email address format'});
        required(schemaPath.password, {message: 'Password is required'});
        required(schemaPath.name, {message: 'Name is required'});
    });

    submit(event: Event): void {
        event.preventDefault();
        this.loading = true;
        this.store.dispatch(
            actionEmployeeCreate({data: this.employeeFormModel()}),
        );
    }
}
