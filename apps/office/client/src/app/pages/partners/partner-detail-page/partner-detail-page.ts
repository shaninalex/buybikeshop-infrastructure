import { Component, inject, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { ActivatedRoute } from '@angular/router';
import { filter, map, Observable, switchMap, tap } from 'rxjs';
import { actionPartnerGet } from '@entities/partner/model/partner.actions';
import { PartnerModel, selectPartnerByID } from '@entities/partner';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-partner-detail-page',
    template: `
        @if (partner$ | async; as partner) {
            <div class="container py-4">
                <div class="d-flex justify-content-between align-items-center mb-4">
                    <div>
                        <h2 class="mb-0">{{ partner.title }}</h2>
                        <small class="text-muted">{{ partner.type }}</small>
                    </div>
                    <div>
                        <button class="btn btn-outline-primary me-2">Edit</button>
                        <button class="btn btn-outline-danger">Delete</button>
                    </div>
                </div>

                <div class="row g-4">
                    <div class="col-md-4">
                        <div class="card">
                            <div class="card-header">
                                General Info
                            </div>
                            <div class="card-body">
                                <p><strong>Status:</strong> {{ partner.active }}</p>
                                <p><strong>Type:</strong> {{ partner.type }}</p>

                                <p><strong>Roles:</strong></p>
                                <div>
                                    @if (partner.is_supplier) {
                                        <span class="badge bg-primary me-1">Supplier</span>
                                    }
                                </div>

                                <hr>

                                <p><strong>Flags:</strong></p>
                                <div>
                                    <span class="badge bg-success me-1">Supplier</span>
                                    <span class="badge bg-warning text-dark me-1">Contractor</span>
                                </div>
                            </div>
                        </div>

                        <div class="card mt-4">
                            <div class="card-header">
                                Bussines Address
                            </div>
                            <div class="card-body">
                                <p class="mb-1">123 Business Street</p>
                                <p class="mb-1">Kyiv</p>
                                <p class="mb-0">Ukraine</p>
                            </div>
                        </div>
                    </div>

                    <div class="col-md-8">
                        <div class="card">
                            <div class="card-header d-flex justify-content-between align-items-center">
                                <span>Contacts</span>
                                <button class="btn btn-sm btn-primary">Add Contact</button>
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
                                        <td>John Doe</td>
                                        <td>Manager</td>
                                        <td>john&#64;acme.com</td>
                                        <td>+380 99 123 4567</td>
                                        <td class="text-end">
                                            <button class="btn btn-sm btn-outline-secondary me-1">Edit</button>
                                            <button class="btn btn-sm btn-outline-danger">Delete</button>
                                        </td>
                                    </tr>

                                    <tr>
                                        <td>Jane Smith</td>
                                        <td>Accountant</td>
                                        <td>jane&#64;acme.com</td>
                                        <td>+380 67 987 6543</td>
                                        <td class="text-end">
                                            <button class="btn btn-sm btn-outline-secondary me-1">Edit</button>
                                            <button class="btn btn-sm btn-outline-danger">Delete</button>
                                        </td>
                                    </tr>
                                    </tbody>
                                </table>
                            </div>
                        </div>

                        <div class="card mt-4">
                            <div class="card-header">Notes</div>

                            <div class="card-body">
                                <div class="mb-4">
                                    <textarea class="form-control" rows="3" placeholder="Add a new note..."></textarea>
                                    <div class="d-flex justify-content-end mt-2">
                                        <button class="btn btn-primary btn-sm">Add Note</button>
                                    </div>
                                </div>

                                <hr>

                                <div class="list-group">
                                    <div class="list-group-item">
                                        <div class="d-flex justify-content-between">
                                            <strong>John Manager</strong>
                                            <small class="text-muted">2026-03-25 10:15</small>
                                        </div>
                                        <p class="mb-2 mt-2">
                                            Discussed new supply contract. Waiting for confirmation.
                                        </p>
                                        <div class="text-end">
                                            <button class="btn btn-sm btn-outline-secondary me-1">Edit</button>
                                            <button class="btn btn-sm btn-outline-danger">Delete</button>
                                        </div>
                                    </div>

                                    <div class="list-group-item">
                                        <div class="d-flex justify-content-between">
                                            <strong>Jane Accountant</strong>
                                            <small class="text-muted">2026-03-20 14:02</small>
                                        </div>
                                        <p class="mb-2 mt-2">
                                            Invoice #456 delayed due to missing documents.
                                        </p>
                                        <div class="text-end">
                                            <button class="btn btn-sm btn-outline-secondary me-1">Edit</button>
                                            <button class="btn btn-sm btn-outline-danger">Delete</button>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        }
    `,
    imports: [
        AsyncPipe
    ]
})
export class PartnerDetailPage implements OnInit {
    private store = inject(Store)
    private route = inject(ActivatedRoute);

    partner$: Observable<PartnerModel> = this.route.paramMap.pipe(
        map(paramsMap => paramsMap.get("id")),
        filter(id => id !== null),
        switchMap(id => this.store.select(selectPartnerByID(parseInt(id))).pipe(
            filter(partner => partner !== undefined)
        ))
    );

    ngOnInit(): void {
        this.route.paramMap.pipe(
            map(paramsMap => paramsMap.get("id")),
            filter(id => id !== null),
            tap(id => this.store.dispatch(actionPartnerGet({ partnerId: parseInt(id) })))
        ).subscribe()
    }
}
