import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { NewProductModel, ProductModel } from '@entities/product';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class ProductApi {
    http = inject(HttpClient);

    GetProducts(): Observable<ProductModel[]> {
        return this.http
            .get<APIResponse<ProductModel[]>>(`/api/v1/products`, {withCredentials: true})
            .pipe(map((response) => response.data));
    }

    CreateProduct(payload: NewProductModel): Observable<ProductModel> {
        return this.http
            .post<APIResponse<ProductModel>>(`/api/v1/products`, payload, {withCredentials: true})
            .pipe(map((response) => response.data));
    }

    DeleteProduct(productId: string): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/product/${productId}`, {withCredentials: true})
            .pipe(map((response) => response.data));
    }

    // NOTE: may be will be better to use UpdateProductModel instead?
    Patch(productId: string, payload: NewProductModel): Observable<ProductModel> {
        return this.http
            .patch<
                APIResponse<ProductModel>
            >(`/api/v1/product/${productId}`, payload, {withCredentials: true})
            .pipe(map((response) => response.data));
    }
}
