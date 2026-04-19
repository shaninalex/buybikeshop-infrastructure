import { createReducer, on } from '@ngrx/store';
import { PartnerRoleModel } from './partner-role.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { actionPartnerRoleSetList } from './partner-role.actions';

export interface PartnerRoleState extends EntityState<PartnerRoleModel> {
}

export const PartnerRolesAdapter = createEntityAdapter<PartnerRoleModel>();
export const initialState = PartnerRolesAdapter.getInitialState();

export const partnerRoleReducer = createReducer(
    initialState,
    on(actionPartnerRoleSetList, (state, action) => PartnerRolesAdapter.addMany(action.PartnerRoles, state)),
);

