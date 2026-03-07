import { Component } from '@angular/core';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'app-sidebar',
    imports: [
        RouterLink
    ],
    host: {
        class: "sidebar"
    },
    template: `
        <aside class="border-end vh-100 p-2 d-flex flex-column justify-content-between">
            <h3 class="fw-bold fs-3">Warehouse</h3>
            <div class="d-flex flex-column justify-content-between flex-grow-1">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" href="#">Inventory</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" routerLink="/vendors">Vendors</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" routerLink="/deliveries">Deliveries</a>
                    </li>
                </ul>
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" routerLink="/support">Support</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" routerLink="/settings">Settings</a>
                    </li>
                </ul>
            </div>
        </aside>
    `,
})
export class SidebarComponent {
}
