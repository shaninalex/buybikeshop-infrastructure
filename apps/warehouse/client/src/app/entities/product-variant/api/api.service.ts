import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient, HttpParams } from '@angular/common/http';
import { ProductVariantModel } from '@entities/product-variant';

@Injectable({
    providedIn: 'root',
})
export class ProductVariantApi {
    http = inject(HttpClient);

    GetProductVariants(productId: number): Observable<ProductVariantModel[]> {
        const params = new HttpParams()
        params.append("productId", productId)
        return this.http
            .get<APIResponse<ProductVariantModel[]>>(`/api/v1/products-variant`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }
}

