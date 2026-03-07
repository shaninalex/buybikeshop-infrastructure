import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'app-inventory-navbar',
    imports: [
        RouterLink
    ],
    template: `
        <nav class="navbar navbar-expand-lg bg-body-tertiary mb-2">
            <div class="container-fluid">
                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item">
                            <a class="nav-link active" routerLink="/">List</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" routerLink="inventory/import">Import</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link disabled">Create</a>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    `,
})
export class InventoryNavbar {

}
