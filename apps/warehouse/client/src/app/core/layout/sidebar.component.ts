import { Component } from '@angular/core';

@Component({
    selector: 'app-sidebar',
    template: `
        <aside class="border-end vh-100 p-2 d-flex flex-column justify-content-between">
            <h3 class="fw-bold fs-3">Warehouse</h3>
            <div class="d-flex flex-column justify-content-between flex-grow-1">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" href="#">Inventory</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Vendors</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Deliveries</a>
                    </li>
                </ul>
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" href="#">Support</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" href="#">Settings</a>
                    </li>
                </ul>
            </div>
        </aside>
    `,
    styleUrl: './sidebar.component.css',
})
export class SidebarComponent {
}
