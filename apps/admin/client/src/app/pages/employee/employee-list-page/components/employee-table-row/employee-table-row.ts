import {Component, Input, OnInit} from '@angular/core';
import { EmployeeModel } from '@entities/employee';
import { RouterLink } from '@angular/router';
import {IdentityTraits} from '@entities/employee/model/employee.model';

@Component({
    selector: 'tr[app-employee-table-row]',
    imports: [RouterLink],
    template: `
        <td>
            <a class="d-flex align-items-center" [routerLink]="['/employee', employee.identity.id]">
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
        <td><span class="badge bg-warning text-dark">Pending</span></td>
        <td class="text-end">
            <a
                class="btn btn-sm btn-outline-secondary"
                [routerLink]="['/employee', employee.identity.id]"
            >
                <i class="fa-solid fa-person"></i>
            </a>
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
