import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import {
    // actionProductCreate,
    // actionProductDelete,
    // actionProductDeleteSuccess,
    actionProductList,
    actionProductSetList,
    // actionProductUpdate,
    // actionProductUpsert,
} from './product.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { ProductApi } from '../api/api.service';

@Injectable()
export class ProductEffects {
    private actions$ = inject(Actions);
    private productsApi = inject(ProductApi);

    get_products_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProductList),
            exhaustMap(() =>
                this.productsApi
                    .GetProducts()
                    .pipe(switchMap((data) => of(actionProductSetList({ products: data })))),
            ),
        ),
    );

    // create_product$ = createEffect(() =>
    //     this.actions$.pipe(
    //         ofType(actionProductCreate),
    //         exhaustMap((action) =>
    //             this.productsApi
    //                 .CreateProduct(action.payload)
    //                 .pipe(switchMap((data) => of(actionProductUpsert({ product: data })))),
    //         ),
    //     ),
    // );

    // update_product$ = createEffect(() =>
    //     this.actions$.pipe(
    //         ofType(actionProductUpdate),
    //         exhaustMap((action) =>
    //             this.productsApi
    //                 .Patch(action.id, action.data)
    //                 .pipe(switchMap((data) => of(actionProductUpsert({ product: data })))),
    //         ),
    //     ),
    // );

    // delete_product$ = createEffect(() =>
    //     this.actions$.pipe(
    //         ofType(actionProductDelete),
    //         exhaustMap((action) =>
    //             this.productsApi
    //                 .DeleteProduct(action.product_id)
    //                 .pipe(
    //                     switchMap(() =>
    //                         of(actionProductDeleteSuccess({ product_id: action.product_id })),
    //                     ),
    //                 ),
    //         ),
    //     ),
    // );
}
