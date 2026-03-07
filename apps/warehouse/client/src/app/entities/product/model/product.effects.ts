import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { actionProductGetList, actionProductSetList, } from './product.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { ProductApi } from '../api/api.service';

@Injectable()
export class ProductEffects {
    private actions$ = inject(Actions);
    private productsApi = inject(ProductApi);

    get_products_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProductGetList),
            exhaustMap(() =>
                this.productsApi
                    .GetProducts()
                    .pipe(switchMap((data) => of(actionProductSetList({ products: data })))),
            ),
        ),
    );
}
