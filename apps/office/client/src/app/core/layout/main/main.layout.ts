import { Component } from '@angular/core';
import { HeaderComponent } from './components';

@Component({
    selector: 'main-layout',
    imports: [
        HeaderComponent
    ],
    template: `
        <app-header/>
        <div class="container">
            <ng-content/>
        </div>
    `
})
export class MainLayout {
}
