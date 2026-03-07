import { ProductEffects, productReducer } from '@entities/product';
import { ProductVariantEffects, productVariantReducer } from '@entities/product-variant';

export const effects = [ProductEffects, ProductVariantEffects];

export const reducers = {
    product: productReducer,
    productVariant: productVariantReducer,
};
