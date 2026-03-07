import { Routes } from '@angular/router';
import { PageContainer } from '@pages/page.container';
import { InventoryPage, inventoryRoutes } from '@pages/inventory-page';
import { VendorsPage } from "@pages/vendors-page";
import { DeliveriesPage } from "@pages/deliveries-page";
import { SettingsPage } from "@pages/settings-page";
import { SupportPage } from "@pages/support-page";

export const routes: Routes = [
    {
        path: '',
        component: PageContainer,
        children: [
            {
                path: '',
                component: InventoryPage,
                children: inventoryRoutes,
            },
            {
                path: 'vendors',
                component: VendorsPage,
            },
            {
                path: 'deliveries',
                component: DeliveriesPage,
            },
            {
                path: 'settings',
                component: SettingsPage,
            },
            {
                path: 'support',
                component: SupportPage,
            },
        ],
    },
];
