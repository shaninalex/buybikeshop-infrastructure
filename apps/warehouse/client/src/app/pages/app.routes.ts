import { Routes } from '@angular/router';
import { HomePage } from './home-page/home-page';
import { PageContainer } from './page.container';

export const routes: Routes = [
    {
        path: '',
        component: PageContainer,
        children: [
            {
                path: '',
                component: HomePage,
            },
        ],
    },
];
