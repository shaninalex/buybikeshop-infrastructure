import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { PartnerRoleModel } from '@entities/partner-role';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { PartnerRolePayloadModel } from '@entities/partner-role/model/partner-role.model';

@Injectable({
    providedIn: 'root',
})
export class PartnerRoleApi {
    http = inject(HttpClient);

    GetPartnerRoles(): Observable<PartnerRoleModel[]> {
        return this.http
            .get<APIResponse<PartnerRoleModel[]>>(`/api/v1/office/partners/roles`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    PatchPartnerRoles(id: number, payload: PartnerRolePayloadModel): Observable<PartnerRoleModel> {
        return this.http
            .patch<APIResponse<PartnerRoleModel>>(`/api/v1/office/partners/roles/${id}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    CreatePartnerRoles(payload: PartnerRolePayloadModel): Observable<PartnerRoleModel> {
        return this.http
            .post<APIResponse<PartnerRoleModel>>(`/api/v1/office/partners/roles`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
