import { Component, inject } from '@angular/core';
import { FormBuilder, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';

@Component({
    selector: 'app-employee-create-google-form',
    imports: [
        FormsModule,
        ReactiveFormsModule
    ],
    template: `
        <form [formGroup]="form" (ngSubmit)="onSubmit()">
            <div class="alert alert-info">
                The employee will be created using your organization's Google
                Workspace account.
            </div>


            <div class="mb-3">
                <label class="form-label">Google Account Email</label>
                <input
                    type="email"
                    class="form-control"
                    formControlName="googleEmail"
                    placeholder="employee@your-company.com"
                />
            </div>

            <button
                type="button"
                class="btn btn-outline-danger"
                disabled
            >
                <i class="bi bi-google"></i>
                Connect Google Account
            </button>

            <div class="alert alert-warning mt-4">Not implemented</div>
        </form>
    `,
})
export class GoogleForm {
    private fb = inject(FormBuilder)

    ngOnInit() {
    }

    form = this.fb.group({
        googleEmail: ['', Validators.required]
    });

    onSubmit(): void {

    }
}
