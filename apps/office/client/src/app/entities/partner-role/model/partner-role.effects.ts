import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import {
    actionPartnerRoleCreate,
    actionPartnerRoleCreateComplete,
    actionPartnerRoleCreateError,
    actionPartnerRoleGetList,
    actionPartnerRolePatch,
    actionPartnerRolePatchComplete,
    actionPartnerRolePatchError,
    actionPartnerRoleSetList,
} from './partner-role.actions';
import { catchError, exhaustMap, map, of, switchMap } from 'rxjs';
import { PartnerRoleApi } from '../api/api.service';
import { HttpErrorResponse } from '@angular/common/http';

@Injectable()
export class PartnerRoleEffects {
    private actions$ = inject(Actions);
    private api = inject(PartnerRoleApi);

    get_roles_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionPartnerRoleGetList),
            exhaustMap(() =>
                this.api
                    .GetPartnerRoles()
                    .pipe(switchMap((roles) => of(actionPartnerRoleSetList({roles})))),
            ),
        ),
    );

    create_role$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionPartnerRoleCreate),
            exhaustMap(action =>
                this.api.CreatePartnerRoles(action.payload).pipe(
                    map(role => actionPartnerRoleCreateComplete({role})),
                    catchError((errors: HttpErrorResponse) => of(actionPartnerRoleCreateError({error: errors.error.errors ?? []}))),
                ),
            ),
        ),
    );

    patch_role$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionPartnerRolePatch),
            exhaustMap(action =>
                this.api.PatchPartnerRoles(action.id, action.payload).pipe(
                    map(role => actionPartnerRolePatchComplete({role})),
                    catchError((errors: HttpErrorResponse) => of(actionPartnerRolePatchError({error: errors.error.errors ?? []}))),
                ),
            ),
        ),
    );
}
