import { Routes } from '@angular/router';
import { PageContainer } from './page.container';
import { employee_routes } from './employee';
import { static_routes } from './static/routes';

export const routes: Routes = [
    {
        path: '',
        component: PageContainer,
        children: [...static_routes, ...employee_routes],
    },
];
