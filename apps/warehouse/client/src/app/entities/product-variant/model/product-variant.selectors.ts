import { createFeatureSelector, createSelector } from '@ngrx/store';
import { productsVariantAdapter, ProductVariantState } from './product-variant.reducer';

export const selectProductsVariantFeature = createFeatureSelector<ProductVariantState>('variant');
export const productsVariantSelectors = productsVariantAdapter.getSelectors();

export const selectProductVariants = createSelector(selectProductsVariantFeature, (state) =>
    productsVariantSelectors.selectAll(state),
);

export const selectProjectVariantByID = (id: number) =>
    createSelector(selectProductsVariantFeature, (state: ProductVariantState) =>
        productsVariantSelectors.selectAll(state).find((p) => p.id === id),
    );
