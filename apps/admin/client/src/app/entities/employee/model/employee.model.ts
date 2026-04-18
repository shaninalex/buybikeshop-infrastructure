import { Identity, IdentityStateEnum } from '@ory/client';

export interface IdentityTraits {
    email: string;
    name: string;
    photo: string;
    dob: Date;
    phone: string;
}

export interface EmployeeModel {
    identity: Identity;
}


export interface EmployeeViewModel {
    id: string;
    email: string;
    name: string;
    photo: string;
    dob: Date;
    phone: string;
    state?: IdentityStateEnum;
    created_at: Date | null;
}

export interface EmployeeCreateFormModel {
    email: string;
    name: string;
    photo: string;
    dob: Date;
    phone: string;
    password?: string;
}
