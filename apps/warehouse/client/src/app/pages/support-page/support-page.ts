import { Component, inject, OnInit } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-support-page',
    imports: [],
    templateUrl: './support-page.html',
})
export class SupportPage implements OnInit {
    private ui = inject(UiService);

    ngOnInit() {
        this.ui.setPageTitle("Support")
    }
}
