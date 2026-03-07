import { Routes } from '@angular/router';
import { InventoryImport, InventoryList } from '@pages/inventory-page/containers';

export const inventoryRoutes: Routes = [
    {
        path: '',
        component: InventoryList
    },
    {
        path: 'inventory/import',
        component: InventoryImport,
    }
]
