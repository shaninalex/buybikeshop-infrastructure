import { createReducer, on } from '@ngrx/store';
import { ProductVariantModel } from './product-variant.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { actionProductVariantSetList } from './product-variant.actions';

export interface ProductVariantState extends EntityState<ProductVariantModel> {
}

export const productsVariantAdapter = createEntityAdapter<ProductVariantModel>();
const initialState = productsVariantAdapter.getInitialState();

export const productVariantReducer = createReducer(
    initialState,
    on(actionProductVariantSetList, (state, action) => productsVariantAdapter.addMany(action.variants, state)),
);

