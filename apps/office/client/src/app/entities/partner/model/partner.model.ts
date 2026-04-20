export enum PartnerType {
    COMPANY = 'company',
    PERSON = 'person',
}

export interface PartnerContact {
    id: number
    contacts: string
    created_at: Date
}

export interface PartnerModel {
    id: number
    title: string
    type: PartnerType
    roles: number[]
    contacts: PartnerContact[]
    created_at: Date
    is_supplier: boolean
    active: boolean
}

export function NewPartnerModel(): PartnerModel {
    return {
        id: 0,
        title: '',
        type: PartnerType.COMPANY,
        roles: [],
        contacts: [],
        created_at: new Date(),
        is_supplier: false,
        active: false,
    }
}
