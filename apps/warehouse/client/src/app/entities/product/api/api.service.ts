import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { ProductModel } from '@entities/product';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class ProductApi {
    http = inject(HttpClient);

    GetProducts(): Observable<ProductModel[]> {
        return this.http
            .get<APIResponse<ProductModel[]>>(`/api/v1/warehouse/products`)
            .pipe(map((response) => response.data));
    }
}
