import { Component, inject, OnInit } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-deliveries-page',
    imports: [],
    templateUrl: './deliveries-page.html',
})
export class DeliveriesPage implements OnInit {
    private ui = inject(UiService);

    ngOnInit() {
        this.ui.setPageTitle("Deliveries")
    }
}
