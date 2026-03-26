import { createAction, props } from '@ngrx/store';
import { ProductModel } from './product.model';

export const actionProductGetList = createAction('[catalog] get list');
export const actionProductSetList = createAction(
    '[Project] set list',
    props<{ products: ProductModel[] }>(),
);
