import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { PartnerRoleModel } from '@entities/partner-role';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class PartnerRoleApi {
    http = inject(HttpClient);

    GetPartnerRoles(): Observable<PartnerRoleModel[]> {
        return this.http
            .get<APIResponse<PartnerRoleModel[]>>(`/api/v1/office/partners/roles`)
            .pipe(map((response) => response.data));
    }
}
