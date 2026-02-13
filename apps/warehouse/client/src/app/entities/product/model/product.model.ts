export interface ProductModel {
    id: number;
    title: string;
    collection_id: number;
    category_id: number;
    brand_id: number;
    supplier_id: number;
    description: string;
    short_description: string;
    created_at: Date;
    updated_at: Date;
}

export interface NewProductModel {
    title: string;
    collection_id: number;
    category_id: number;
    brand_id: number;
    supplier_id: number;
    description: string;
    short_description: string;
}

export interface ProductVariantModel {
    id: string;
    product_id: number;
    title: string;
    description: string;
    sku: string;
    barcode: string;
    created_at: Date;
    updated_at: Date;
}

export interface NewProductVariantModel {
    product_id: number;
    title: string;
    description: string;
    sku: string;
    barcode: string;
}
