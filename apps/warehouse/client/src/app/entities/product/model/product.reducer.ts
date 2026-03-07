import { createReducer, on } from '@ngrx/store';
import { ProductModel } from './product.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { actionProductSetList } from './product.actions';

export interface ProductState extends EntityState<ProductModel> {
}

export const productsAdapter = createEntityAdapter<ProductModel>();
export const initialState = productsAdapter.getInitialState();

export const productReducer = createReducer(
    initialState,
    on(actionProductSetList, (state, action) => productsAdapter.addMany(action.products, state)),
);

