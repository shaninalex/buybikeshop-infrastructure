import { createAction, props } from '@ngrx/store';
import { PartnerRoleModel } from './partner-role.model';

export const actionPartnerRoleGetList = createAction('[PartnerRole] get list');
export const actionPartnerRoleSetList = createAction(
    '[PartnerRole] set list',
    props<{ PartnerRoles: PartnerRoleModel[] }>(),
);
