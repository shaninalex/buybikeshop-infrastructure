import { Component, inject, signal } from '@angular/core';
import { PartnerRoleSelector } from '@entities/partner-role';
import { form, FormField, required } from '@angular/forms/signals';
import { NewPartnerModel, PartnerModel } from '@entities/partner';
import { FormsModule } from '@angular/forms';
import { Store } from '@ngrx/store';
import { actionPartnerCreate } from '@entities/partner/model/partner.actions';
import { PartnerType } from '@entities/partner/model/partner.model';

@Component({
    selector: 'app-partners-create-page',
    imports: [
        PartnerRoleSelector,
        FormField,
        FormsModule
    ],
    template: `
        <form (submit)="submit($event)" class="container py-4">
            <div class="d-flex justify-content-between align-items-center mb-4">
                <div>
                    <h2 class="mb-0">Create New Partner</h2>
                    <small class="text-muted">Company or Person</small>
                </div>
                <div class="btn-group">
                    <button class="btn btn-outline-secondary">Cancel</button>
                    <button class="btn btn-primary">Save Partner</button>
                </div>
            </div>

            <div class="row g-4">
                <div class="col-md-4">
                    <div class="card">
                        <div class="card-header">General Info</div>
                        <div class="card-body">
                            <div class="mb-3">
                                <label class="form-label">Name</label>
                                <input type="text"
                                       class="form-control"
                                       placeholder="Enter name"
                                       [formField]="partnerForm.title"
                                />
                            </div>
                            <div class="mb-3">
                                <label class="form-label">Type</label>
                                <select class="form-select" [formField]="partnerForm.type">
                                    <option selected [value]="PartnerType.COMPANY">Company</option>
                                    <option [value]="PartnerType.PERSON">Person</option>
                                </select>
                            </div>
                            <div class="mb-3">
                                <div class="form-check">
                                    <input class="form-check-input" type="checkbox" id="status"
                                           [formField]="partnerForm.active">
                                    <label class="form-check-label" for="status">
                                        Status (active/inactive)
                                    </label>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="card mt-4">
                        <div class="card-header">Roles</div>
                        <div class="card-body">
                            <app-partner-role-selector [formField]="partnerForm.roles"/>
                        </div>
                    </div>

                    <div class="card mt-4">
                        <div class="card-header">Flags</div>
                        <div class="card-body">

                            <div class="form-check">
                                <input class="form-check-input" type="checkbox" id="flagSupplier"
                                       [formField]="partnerForm.is_supplier">
                                <label class="form-check-label" for="flagSupplier">
                                    Supplier Enabled
                                </label>
                            </div>

                            <div class="form-check">
                                <input class="form-check-input" type="checkbox" id="flagContractor">
                                <label class="form-check-label" for="flagContractor">
                                    Contractor
                                </label>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="col-md-8">
                    <div class="card">
                        <div class="card-header">Address</div>
                        <div class="card-body">

                            <div class="mb-3">
                                <label class="form-label">Street</label>
                                <input type="text" class="form-control" placeholder="Street address">
                            </div>

                            <div class="row">
                                <div class="col-md-6 mb-3">
                                    <label class="form-label">City</label>
                                    <input type="text" class="form-control" placeholder="City">
                                </div>

                                <div class="col-md-6 mb-3">
                                    <label class="form-label">Country</label>
                                    <input type="text" class="form-control" placeholder="Country">
                                </div>
                            </div>

                        </div>
                    </div>

                    <div class="card mt-4">
                        <div class="card-header d-flex justify-content-between align-items-center">
                            <span>Contacts</span>
                            <button class="btn btn-sm btn-outline-primary">Add Contact</button>
                        </div>

                        <div class="card-body p-0">
                            <table class="table mb-0">
                                <thead class="table-light">
                                <tr>
                                    <th>Name</th>
                                    <th>Role</th>
                                    <th>Email</th>
                                    <th>Phone</th>
                                    <th></th>
                                </tr>
                                </thead>
                                <tbody>
                                <tr>
                                    <td>—</td>
                                    <td>—</td>
                                    <td>—</td>
                                    <td>—</td>
                                    <td class="text-end text-muted">No contacts yet</td>
                                </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>

                    <div class="card mt-4">
                        <div class="card-header">Notes</div>
                        <div class="card-body">
                            <textarea disabled class="form-control" rows="3"
                                      placeholder="Add initial note..."></textarea>
                            <div class="d-flex justify-content-end mt-2">
                                <button disabled class="btn btn-primary btn-sm">Add Note</button>
                            </div>
                            <div class="text-muted mt-3">
                                No notes yet
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </form>
    `,
})
export class PartnersCreatePage {
    private store = inject(Store);
    partnerFormModel = signal<PartnerModel>(NewPartnerModel());

    partnerForm = form(this.partnerFormModel, (schemaPath) => {
        required(schemaPath.title, {message: "Role name is required"});
    });

    submit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionPartnerCreate({payload: this.partnerFormModel()}));
    }

    protected readonly PartnerType = PartnerType;
}




