import { Component, inject, OnInit } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-home-page',
    imports: [],
    templateUrl: './home-page.html',
})
export class HomePage implements OnInit {
    private ui = inject(UiService);

    ngOnInit() {
        this.ui.setPageTitle("Inventory")
    }
}
