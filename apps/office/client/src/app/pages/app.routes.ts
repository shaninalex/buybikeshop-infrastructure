import { Routes } from '@angular/router';
import { HomePage } from '@pages/home';
import { PageContainer } from '@pages/page.container';
import { partners_routes } from "@pages/partners/routes";
import { catalog_routes } from "@pages/catalog";

export const routes: Routes = [
    {
        path: '',
        component: PageContainer,
        children: [
            {
                path: "",
                component: HomePage,
            },
            ...partners_routes,
            ...catalog_routes,
        ]
    }
];
