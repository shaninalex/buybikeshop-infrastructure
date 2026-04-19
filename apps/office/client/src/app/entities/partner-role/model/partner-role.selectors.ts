import { createFeatureSelector, createSelector } from '@ngrx/store';
import { PartnerRolesAdapter, PartnerRoleState } from './partner-role.reducer';

export const selectPartnerRolesFeature = createFeatureSelector<PartnerRoleState>('partner_role');
export const PartnerRolesSelectors = PartnerRolesAdapter.getSelectors();

export const selectPartnerRoles = createSelector(selectPartnerRolesFeature, (state) =>
    PartnerRolesSelectors.selectAll(state),
);

export const selectPartnerRoleByID = (id: number) =>
    createSelector(selectPartnerRolesFeature, (state: PartnerRoleState) =>
        PartnerRolesSelectors.selectAll(state).find((p) => p.id === id),
    );
