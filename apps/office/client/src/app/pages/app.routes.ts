import { Routes } from '@angular/router';
import { HomePage } from '@pages/home';
import { PageContainer } from '@pages/page.container';
import { PartnersPage } from '@pages/partners';

export const routes: Routes = [
    {
        path: '',
        component: PageContainer,
        children: [
            {
                path: "",
                component: HomePage,
            },
            {
                path: "partners",
                component: PartnersPage,
            }
        ]
    }
];
