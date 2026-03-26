import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { PartnerModel } from '@entities/partner';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

interface partnersResponse {
    partners: PartnerModel[]
}

@Injectable({
    providedIn: 'root',
})
export class PartnerApi {
    http = inject(HttpClient);

    GetPartners(): Observable<PartnerModel[]> {
        return this.http
            .get<APIResponse<partnersResponse>>(`/api/v1/office/partners`)
            .pipe(
                map((response) => response.data),
                map((response) => response.partners),
            );
    }
}
