import { Routes } from '@angular/router';
import { PageContainer } from './page.container';
import { HomePage } from './static';

export const routes: Routes = [
    {
        path: '',
        component: PageContainer,
        children: [
            {
                path: '',
                component: HomePage,
            },
        ]
    }
];
