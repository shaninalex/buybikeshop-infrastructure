import { createAction, props } from '@ngrx/store';
import { ProductVariantModel } from './product-variant.model';

export const actionProductVariantGetList = createAction(
    '[Project Variant] get list',
    props<{ productId: number }>(),
);
export const actionProductVariantSetList = createAction(
    '[Project Variant] set list',
    props<{ variants: ProductVariantModel[] }>(),
);
