import { Component, inject, OnInit } from '@angular/core';
import { RouterLink } from "@angular/router";
import { Store } from '@ngrx/store';
import { actionPartnerGetList, selectPartners } from '@entities/partner';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-partners-list-page',
    imports: [
        RouterLink,
        AsyncPipe
    ],
    template: `
        <h3>Partners page</h3>

        <nav class="navbar bg-body-tertiary mb-4">
            <div class="container-fluid">
                <form class="d-flex" role="search">
                    <input class="form-control me-2" type="search" placeholder="Search partner" aria-label="Search"/>
                    <button class="btn btn-outline-success" type="submit">Search</button>
                </form>

                <a routerLink="/partners/create" class="btn btn-primary">Create</a>
            </div>
        </nav>

        <table class="table">
            <thead>
            <tr>
                <th scope="col">Type</th>
                <th scope="col" style="width: 100%">Name</th>
                <th scope="col">Supplier</th>
            </tr>
            </thead>
            <tbody>
                @if (partners$ | async; as partners) {
                    @for (partner of partners; track $index) {
                        <tr>
                            <th scope="row">
                                @if (partner.type === 'company') {
                                    <i class="fa-regular fa-building"></i>
                                } @else {
                                    <i class="fa-solid fa-person"></i>
                                }
                            </th>
                            <td>
                                <a [routerLink]='["/partners", partner.id]' class="text-body">{{ partner.title }}</a>
                            </td>
                            <th scope="row">
                                @if (partner.is_supplier) {
                                    <i class="fa-solid fa-truck-field"></i>
                                } @else {
                                    <span>-</span>
                                }
                            </th>
                        </tr>
                    }
                }
            </tbody>
        </table>
    `,
})
export class PartnersListPage implements OnInit {
    private store = inject(Store);
    partners$ = this.store.select(selectPartners)

    ngOnInit(): void {
        this.store.dispatch(actionPartnerGetList())
    }
}
