import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { actionPartnerGet, actionPartnerGetList, actionPartnerSet, actionPartnerSetList, } from './partner.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { PartnerApi } from '../api/api.service';

@Injectable()
export class PartnerEffects {
    private actions$ = inject(Actions);
    private partnersApi = inject(PartnerApi);

    get_partners_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionPartnerGetList),
            exhaustMap(() =>
                this.partnersApi
                    .GetPartners()
                    .pipe(switchMap((data) => of(actionPartnerSetList({ partners: data })))),
            ),
        ),
    );

    get_partner$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionPartnerGet),
            exhaustMap((action) =>
                this.partnersApi
                    .GetPartner(action.partnerId)
                    .pipe(switchMap((data) => of(actionPartnerSet({ partner: data })))),
            ),
        ),
    );
}
