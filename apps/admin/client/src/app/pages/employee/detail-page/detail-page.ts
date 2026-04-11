import { Component, inject } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { UiService } from '@shared/ui';
import { filter, map, Observable, tap } from 'rxjs';
import { AsyncPipe, DatePipe } from '@angular/common';
import { EmployeeViewModel } from '@entities/employee/model/employee.model';

@Component({
    selector: 'app-employee-detail-page',
    imports: [
        AsyncPipe,
        DatePipe,
        RouterLink,
    ],
    templateUrl: './detail-page.html',
})
export class DetailPage {
    private route = inject(ActivatedRoute);
    private ui = inject(UiService);

    employee$: Observable<EmployeeViewModel> = this.route.data.pipe(
        filter((data) => !!data['employee']),
        map((data) => data['employee'] as EmployeeViewModel),
        tap((employee) => this.ui.setPageTitle(`Employee: ${employee.name}`)),
    );
}
