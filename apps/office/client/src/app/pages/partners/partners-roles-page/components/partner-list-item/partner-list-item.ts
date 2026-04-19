import { Component, DestroyRef, inject, Input, OnInit, signal } from '@angular/core';
import { PartnerRoleModel } from '@entities/partner-role'
import { PartnerRolePayloadModel } from '@entities/partner-role/model/partner-role.model';
import { form, FormField, required } from '@angular/forms/signals';
import { FormInputError } from '@shared/ui';
import { FormsModule } from '@angular/forms';
import {
    actionPartnerRolePatch,
    actionPartnerRolePatchComplete
} from '@entities/partner-role/model/partner-role.actions';
import { Store } from '@ngrx/store';
import { Actions, ofType } from '@ngrx/effects';
import { finalize, tap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Component({
    selector: 'li[app-partner-list-item]',
    imports: [
        FormInputError,
        FormsModule,
        FormField
    ],
    template: `
        @if (isEdit()) {
            <form (submit)="submit($event)" class="d-flex align-items-center gap-2">
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
                    <button type="submit" class="btn btn-outline-success">Save</button>
                </div>
            </form>
        } @else {
            <button class="bg-transparent border-0 p-0" (click)="edit()">
                id:{{ role.id }} - <b>{{ role.role }}</b>
            </button>
        }
    `,
})
export class PartnerListItem implements OnInit {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    @Input() role: PartnerRoleModel;

    isEdit = signal(false);
    roleFormModel = signal<PartnerRolePayloadModel>({
        role: '',
    });

    roleForm = form(this.roleFormModel, (schemaPath) => {
        required(schemaPath.role, {message: "Role name is required"});
    });

    ngOnInit() {
        this.roleFormModel.set({
            id: this.role.id,
            role: this.role.role,
        });
        this.actions$.pipe(
            ofType(actionPartnerRolePatchComplete),
            takeUntilDestroyed(this.destroyRef),
        ).subscribe(() => this.isEdit.set(false));
    }

    submit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(
            actionPartnerRolePatch({id: this.role.id, payload: this.roleFormModel()})
        );
    }

    cancel(): void {
        this.roleForm().reset();
        this.isEdit.set(false);
    }

    edit(): void {
        this.isEdit.set(true);
    }
}
