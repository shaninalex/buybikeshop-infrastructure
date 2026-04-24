import { createAction, props } from '@ngrx/store';
import { PartnerModel } from './partner.model';
import { PartnerRolePayloadModel } from '@entities/partner-role/model/partner-role.model';
import { PartnerRoleModel } from '@entities/partner-role';

export const actionPartnerGetList = createAction('[Partner] get list');

export const actionPartnerSetList = createAction(
    '[Partner] set list',
    props<{ partners: PartnerModel[] }>(),
);

export const actionPartnerGet = createAction(
    '[Partner] get',
    props<{ partnerId: number }>(),
);

export const actionPartnerSet = createAction(
    '[Partner] set',
    props<{ partner: PartnerModel }>(),
);

export const actionPartnerCreate = createAction(
    '[Partner] create',
    props<{ payload: PartnerModel }>(),
);

export const actionPartnerCreateComplete = createAction(
    '[Partner] create complete',
    props<{ partner: PartnerModel }>(),
);

export const actionPartnerCreateError = createAction(
    '[Partner] create error',
    props<{ errors: any }>(),
);
