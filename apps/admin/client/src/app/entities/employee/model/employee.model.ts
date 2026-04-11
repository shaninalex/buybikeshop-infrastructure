import { Identity } from '@ory/client';

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
