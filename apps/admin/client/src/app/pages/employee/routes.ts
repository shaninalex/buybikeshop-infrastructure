import { Routes } from '@angular/router';
import { EmployeeContainer } from './employee.container';
import { EmployeeListPage } from './employee-list-page';

export const employee_routes: Routes = [
    {
        path: 'employee',
        component: EmployeeContainer,
        children: [
            {
                path: '',
                component: EmployeeListPage,
            },
        ],
    },
];
