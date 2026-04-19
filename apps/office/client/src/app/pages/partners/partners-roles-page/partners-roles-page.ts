import { Component, inject, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionPartnerRoleGetList, PartnerRoleModel, selectPartnerRoles } from '@entities/partner-role';
import { AsyncPipe } from '@angular/common';
import { Observable } from 'rxjs';

@Component({
    selector: 'app-partners-roles-page',
    imports: [
        AsyncPipe
    ],
    template: `
        <ul class="list-group">
            @if (roles$ | async; as rr) {
                @for (r of rr; track r.id) {
                    <li class="list-group-item">id:{{ r.id }} - <b>{{ r.role }}</b></li>
                }
            }
            <button class="list-group-item list-group-item-action">
                <i class="fa-regular fa-square-plus"></i>
            </button>
        </ul>
    `,
})
export class PartnersRolesPage implements OnInit {
    private store = inject(Store);
    roles$: Observable<PartnerRoleModel[]> = this.store.select(selectPartnerRoles)

    ngOnInit() {
        this.store.dispatch(actionPartnerRoleGetList())
    }
}
