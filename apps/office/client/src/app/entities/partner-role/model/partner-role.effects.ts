import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { actionPartnerRoleGetList, actionPartnerRoleSetList, } from './partner-role.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { PartnerRoleApi } from '../api/api.service';

@Injectable()
export class PartnerRoleEffects {
    private actions$ = inject(Actions);
    private PartnerRolesApi = inject(PartnerRoleApi);

    get_PartnerRoles_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionPartnerRoleGetList),
            exhaustMap(() =>
                this.PartnerRolesApi
                    .GetPartnerRoles()
                    .pipe(switchMap((data) => of(actionPartnerRoleSetList({ PartnerRoles: data })))),
            ),
        ),
    );
}
