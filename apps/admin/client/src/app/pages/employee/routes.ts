import { Routes } from '@angular/router';

import { EmployeeContainer } from './employee.container';
import { ListPage } from './list-page';
import { DetailPage } from './detail-page';
import { CreatePage } from './create-page';

import { employeeViewModelResolver } from './employee-view-model-resolver';

export const employee_routes: Routes = [
    {
        path: 'employees',
        component: EmployeeContainer,
        children: [
            {
                path: '',
                component: ListPage,
            },
            {
                path: 'create',
                component: CreatePage,
            },
            {
                path: ':id',
                component: DetailPage,
                resolve: {
                    employee: employeeViewModelResolver,
                }
            },
        ],
    },
];
