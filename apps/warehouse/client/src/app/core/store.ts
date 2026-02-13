import { productReducer, ProductEffects } from '@entities/product';

export const effects = [ProductEffects];

export const reducers = {
    product: productReducer,
};
