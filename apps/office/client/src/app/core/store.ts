import { PartnerEffects, partnerReducer } from '@entities/partner';

export const effects = [PartnerEffects];

export const reducers = {
    partner: partnerReducer,
};
