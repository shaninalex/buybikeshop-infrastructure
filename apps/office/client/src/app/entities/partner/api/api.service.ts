import { inject, Injectable } from '@angular/core';
import { filter, map, Observable } from 'rxjs';
import { PartnerModel } from '@entities/partner';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class PartnerApi {
    http = inject(HttpClient);

    GetPartners(): Observable<PartnerModel[]> {
        return this.http
            .get<APIResponse<PartnerModel[]>>(`/api/v1/office/partners`, { withCredentials: true })
            .pipe(
                map((response) => response.data),
                filter((partners) => partners !== null),
            );
    }

    GetPartner(partnerId: number): Observable<PartnerModel> {
        return this.http
            .get<APIResponse<PartnerModel>>(`/api/v1/office/partners/${partnerId}`, { withCredentials: true })
            .pipe(
                map((response) => response.data),
                filter((partner) => partner !== null),
            );
    }
}
