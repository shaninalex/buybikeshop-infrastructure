import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'app-sidebar',
    imports: [RouterLink],
    host: {
        class: 'sidebar',
    },
    template: `
        <aside class="border-end vh-100 p-2 d-flex flex-column justify-content-between">
            <h4 class="fw-bold fs-4">
                <img src="/images/logo.png" alt="" style="width: 44px;" />
                Admin
            </h4>
            <div class="d-flex flex-column justify-content-between flex-grow-1">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" routerLink="/employees">
                            Employees
                        </a>
                    </li>
                </ul>
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" routerLink="/">Settings</a>
                    </li>
                </ul>
            </div>
        </aside>
    `,
})
export class SidebarComponent {}
