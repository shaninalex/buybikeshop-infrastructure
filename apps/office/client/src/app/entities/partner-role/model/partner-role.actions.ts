import { createAction, props } from '@ngrx/store';
import { PartnerRoleModel, PartnerRolePayloadModel } from './partner-role.model';

export const actionPartnerRoleGetList = createAction('[PartnerRole] get list');
export const actionPartnerRoleSetList = createAction(
    '[PartnerRole] set list',
    props<{ roles: PartnerRoleModel[] }>(),
);

export const actionPartnerRolePatch = createAction(
    '[PartnerRole] patch',
    props<{ id: number, payload: PartnerRolePayloadModel }>(),
);
export const actionPartnerRolePatchComplete = createAction(
    '[PartnerRole] patch complete',
    props<{ role: PartnerRoleModel }>(),
);
export const actionPartnerRolePatchError = createAction(
    '[PartnerRole] patch error',
    props<{ error: any }>(),
);

export const actionPartnerRoleCreate = createAction(
    '[PartnerRole] create',
    props<{ payload: PartnerRolePayloadModel }>(),
);
export const actionPartnerRoleCreateComplete = createAction(
    '[PartnerRole] create complete',
    props<{ role: PartnerRoleModel }>(),
);
export const actionPartnerRoleCreateError = createAction(
    '[PartnerRole] create error',
    props<{ error: any }>(),
);
