import { Component, inject, OnInit } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-vendors-page',
    imports: [],
    templateUrl: './vendors-page.html',
})
export class VendorsPage implements OnInit {
    private ui = inject(UiService);

    ngOnInit() {
        this.ui.setPageTitle("Vendors")
    }
}
