import { createAction, props } from '@ngrx/store';
import { PartnerModel } from './partner.model';

export const actionPartnerGetList = createAction('[Partner] get list');
export const actionPartnerSetList = createAction(
    '[Partner] set list',
    props<{ partners: PartnerModel[] }>(),
);
