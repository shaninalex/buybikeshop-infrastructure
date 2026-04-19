import { Component, DestroyRef, inject, OnInit, signal } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionPartnerRoleGetList, selectPartnerRoles } from '@entities/partner-role';
import { AsyncPipe } from '@angular/common';
import { PartnerListItem } from '@pages/partners/partners-roles-page/components';
import {
    actionPartnerRoleCreate,
    actionPartnerRoleCreateComplete,
} from '@entities/partner-role/model/partner-role.actions';
import { PartnerRolePayloadModel } from '@entities/partner-role/model/partner-role.model';
import { form, FormField, required } from '@angular/forms/signals';
import { FormInputError } from '@shared/ui';
import { FormsModule } from '@angular/forms';
import { Actions, ofType } from '@ngrx/effects';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Component({
    selector: 'app-partners-roles-page',
    imports: [
        AsyncPipe,
        PartnerListItem,
        FormInputError,
        FormsModule,
        FormField
    ],
    template: `
        <ul class="list-group">
            @if (roles$ | async; as rr) {
                @for (r of rr; track r.id) {
                    <li class="list-group-item list-group-item-action" app-partner-list-item [role]="r"></li>
                }
            }
            @if (isCreating()) {
                <form (submit)="submit($event)" class="list-group-item d-flex align-items-center gap-2">
                    <div>
                        <input
                            type="text"
                            id="name"
                            class="form-control"
                            [formField]="roleForm.role"
                            placeholder="Enter a name"
                        />
                        <app-form-input-error [inputField]="roleForm.role"/>
                    </div>
                    <div class="btn-group">
                        <button type="button" class="btn btn-outline-secondary" (click)="cancel()">Cancel</button>
                        <button type="submit" class="btn btn-outline-success">Create</button>
                    </div>
                </form>
            } @else {
                <button class="list-group-item list-group-item-action" (click)="create()">
                    <i class="fa-regular fa-square-plus"></i>
                </button>
            }
        </ul>
    `,
})
export class PartnersRolesPage implements OnInit {
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private store = inject(Store);
    roles$ = this.store.select(selectPartnerRoles)
    isCreating = signal(false);
    roleFormModel = signal<PartnerRolePayloadModel>({
        role: '',
    });

    roleForm = form(this.roleFormModel, (schemaPath) => {
        required(schemaPath.role, {message: "Role name is required"});
    });

    ngOnInit() {
        this.store.dispatch(actionPartnerRoleGetList())
        this.actions$.pipe(
            ofType(actionPartnerRoleCreateComplete),
            takeUntilDestroyed(this.destroyRef),
        ).subscribe(() => this.cancel());
    }

    submit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(
            actionPartnerRoleCreate({payload: this.roleFormModel()})
        );
    }

    cancel(): void {
        this.roleForm().reset();
        this.roleFormModel.set({role: ''});
        this.isCreating.set(false);
    }

    create(): void {
        this.isCreating.set(true);
    }
}
