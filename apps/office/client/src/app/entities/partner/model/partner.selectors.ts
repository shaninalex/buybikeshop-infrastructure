import { createFeatureSelector, createSelector } from '@ngrx/store';
import { partnersAdapter, PartnerState } from './partner.reducer';

export const selectPartnersFeature = createFeatureSelector<PartnerState>('partner');
export const partnersSelectors = partnersAdapter.getSelectors();

export const selectPartners = createSelector(selectPartnersFeature, (state) =>
    partnersSelectors.selectAll(state),
);

export const selectPartnerByID = (id: number) =>
    createSelector(selectPartnersFeature, (state: PartnerState) =>
        partnersSelectors.selectAll(state).find((p) => p.id === id),
    );
