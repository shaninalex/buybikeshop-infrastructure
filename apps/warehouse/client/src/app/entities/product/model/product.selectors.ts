import { createFeatureSelector, createSelector } from '@ngrx/store';
import { productsAdapter, ProductState } from './product.reducer';

export const selectProductsFeature = createFeatureSelector<ProductState>('product');
export const productsSelectors = productsAdapter.getSelectors();

export const selectProducts = createSelector(selectProductsFeature, (state) =>
    productsSelectors.selectAll(state),
);

export const selectProjectByID = (id: number) =>
    createSelector(selectProductsFeature, (state: ProductState) =>
        productsSelectors.selectAll(state).find((p) => p.id === id),
    );
