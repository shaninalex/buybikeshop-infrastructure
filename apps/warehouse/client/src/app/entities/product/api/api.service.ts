import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { ProductModel } from '@entities/product';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

interface productsResponse {
    products: ProductModel[]
}

@Injectable({
    providedIn: 'root',
})
export class ProductApi {
    http = inject(HttpClient);

    GetProducts(): Observable<ProductModel[]> {
        return this.http
            .get<APIResponse<productsResponse>>(`/api/v1/warehouse/products`)
            .pipe(
                map((response) => response.data),
                map((response) => response.products),
            );
    }
}
