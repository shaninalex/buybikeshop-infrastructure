import { PartnerEffects, partnerReducer } from '@entities/partner';
import { PartnerRoleEffects, partnerRoleReducer } from '@entities/partner-role';

export const effects = [PartnerEffects, PartnerRoleEffects];

export const reducers = {
    partner: partnerReducer,
    partner_role: partnerRoleReducer,
};
