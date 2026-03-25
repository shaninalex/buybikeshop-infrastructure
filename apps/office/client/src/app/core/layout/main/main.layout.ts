import { Component } from '@angular/core';
import { HeaderComponent } from './components';

@Component({
    selector: 'main-layout',
    imports: [
        HeaderComponent
    ],
    template: `
        <app-header/>
        <div class="container-fluid">
            <ng-content/>
        </div>

    `
})
export class MainLayout {
}
