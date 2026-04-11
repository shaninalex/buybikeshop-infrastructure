import { Component, inject, OnInit } from '@angular/core';
import { actionEmployeeGetList, EmployeeModel, selectEmployees } from '@entities/employee';
import { Store } from '@ngrx/store';
import { EmployeeTableRow } from './components';
import { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-employee-list-page',
    imports: [EmployeeTableRow, AsyncPipe],
    templateUrl: './employee-list-page.html',
})
export class EmployeeListPage implements OnInit {
    private store = inject(Store);
    employees: Observable<EmployeeModel[]> = this.store.select(selectEmployees);

    ngOnInit(): void {
        this.store.dispatch(actionEmployeeGetList());
    }
}
