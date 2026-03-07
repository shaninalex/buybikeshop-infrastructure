import { Component, inject, OnInit } from '@angular/core';
import { UiService } from '@shared/ui';
import { InventoryNavbar } from '@pages/inventory-page/components';
import { RouterOutlet } from '@angular/router';

@Component({
    selector: 'app-inventory-page',
    imports: [
        InventoryNavbar,
        RouterOutlet
    ],
    templateUrl: './inventory-page.html',
})
export class InventoryPage implements OnInit {
    private ui = inject(UiService);

    ngOnInit() {
        this.ui.setPageTitle("Inventory")
    }
}
