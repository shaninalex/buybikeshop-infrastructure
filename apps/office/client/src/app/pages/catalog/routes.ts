import { Routes } from "@angular/router";
import { CatalogContainer } from "@pages/catalog/catalog.container";
import { CatalogMainPage } from "@pages/catalog/catalog-main-page";
import { PurchaseOrderPage } from "@pages/catalog/purchase-order-page";

export const catalog_routes: Routes = [
    {
        path: "catalog",
        component: CatalogContainer,
        children: [
            {
                path: "",
                component: CatalogMainPage,
            },
            {
                path: "purchase-order",
                component: PurchaseOrderPage,
            }
        ]
    }
];
