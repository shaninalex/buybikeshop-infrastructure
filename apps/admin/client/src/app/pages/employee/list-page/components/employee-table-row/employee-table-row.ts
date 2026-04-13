import {Component, Input, OnInit} from '@angular/core';
import { EmployeeModel } from '@entities/employee';
import { RouterLink } from '@angular/router';
import {IdentityTraits} from '@entities/employee/model/employee.model';

@Component({
    selector: 'tr[app-employee-table-row]',
    imports: [RouterLink],
    template: `
        <td>
            <a class="d-flex align-items-center" [routerLink]="['/employees', employee.identity.id]">
                <img
                    [src]="traits.photo"
                    class="rounded-circle me-2"
                    alt="Avatar"
                    style="width: 30px"
                />
                {{ traits.name }}
            </a>
        </td>
        <td>{{ traits.email }}</td>
        <td><span class="badge bg-info text-dark">User</span></td>
        <td>
            @if (employee.identity.state === "active") {
                <span class="badge bg-warning text-dark">{{ employee.identity.state }}</span>
            } @else if (employee.identity.state === "inactive") {
                <span class="badge bg-warning text-dark">{{ employee.identity.state }}</span>
            }
        </td>
    `,
})
export class EmployeeTableRow implements OnInit {
    @Input() employee: EmployeeModel;
    traits: IdentityTraits;

    ngOnInit() {
        this.traits = this.employee.identity.traits;
    }
}
