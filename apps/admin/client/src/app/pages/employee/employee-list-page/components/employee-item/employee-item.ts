import { Component, Input } from '@angular/core';
import { EmployeeModel } from '@entities/employee';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'tr[app-employee-item]',
    imports: [RouterLink],
    template: `
        <td>3</td>
        <td>
            <a class="d-flex align-items-center" [routerLink]="['/employee', employee.identity.id]">
                <img
                    src="/images/default-avatar.png"
                    class="rounded-circle me-2"
                    alt="Avatar"
                    style="width: 30px"
                />
                Michael Brown
            </a>
        </td>
        <td>{{ employee.identity.traits['email'] }}</td>
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
export class EmployeeItem {
    @Input() employee: EmployeeModel;
}
