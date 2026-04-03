export interface PartnerRole {
    id: number;
    role: string
}

export interface PartnerContact {
    id: number;
    contacts: string;
    created_at: Date;
}

export interface PartnerModel {
    id: number;
    role: PartnerRole;
    title: string;
    type: string;
    active: boolean;
    is_supplier: boolean;
    contacts: PartnerContact[];
    created_at: Date;
}
