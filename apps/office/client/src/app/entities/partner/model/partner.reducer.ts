import { createReducer, on } from '@ngrx/store';
import { PartnerModel } from './partner.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { actionPartnerSetList } from './partner.actions';

export interface PartnerState extends EntityState<PartnerModel> {
}

export const partnersAdapter = createEntityAdapter<PartnerModel>();
export const initialState = partnersAdapter.getInitialState();

export const partnerReducer = createReducer(
    initialState,
    on(actionPartnerSetList, (state, action) => partnersAdapter.addMany(action.partners, state)),
);

