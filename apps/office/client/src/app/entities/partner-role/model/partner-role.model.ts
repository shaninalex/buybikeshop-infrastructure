export interface PartnerRoleModel {
    id: number;
    role: string
}

// for patch or create
export interface PartnerRolePayloadModel {
    id?: number;
    role: string
}
