import { createAction, props } from '@ngrx/store';
import { ProductModel } from './product.model';

export const actionProductList = createAction('[product] get list');
export const actionProductSetList = createAction(
    '[Project] set list',
    props<{ products: ProductModel[] }>(),
);
