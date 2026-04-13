import { Component, inject, OnInit } from '@angular/core';
import { actionEmployeeGetList, EmployeeModel, selectEmployees } from '@entities/employee';
import { Store } from '@ngrx/store';
import { EmployeeTableRow } from './components';
import { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { UiService } from '@shared/ui';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'app-employee-list-page',
    imports: [EmployeeTableRow, AsyncPipe, RouterLink],
    templateUrl: './list-page.html',
})
export class ListPage implements OnInit {
    private store = inject(Store);
    private ui = inject(UiService);

    employees: Observable<EmployeeModel[]> = this.store.select(selectEmployees);

    ngOnInit(): void {
        this.store.dispatch(actionEmployeeGetList());
        this.ui.setPageTitle(`Employees`)
    }
}
