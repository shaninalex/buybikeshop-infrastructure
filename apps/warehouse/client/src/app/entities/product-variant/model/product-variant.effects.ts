import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { actionProductVariantGetList, actionProductVariantSetList } from './product-variant.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { ProductVariantApi } from '../api/api.service';

@Injectable()
export class ProductVariantEffects {
    private actions$ = inject(Actions);
    private productVariantApi = inject(ProductVariantApi);

    get_products_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionProductVariantGetList),
            exhaustMap((action) =>
                this.productVariantApi
                    .GetProductVariants(action.productId)
                    .pipe(switchMap((variants) => of(actionProductVariantSetList({ variants })))),
            ),
        ),
    );
}
