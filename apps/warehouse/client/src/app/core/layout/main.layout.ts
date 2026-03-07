import { Component } from '@angular/core';
import { HeaderComponent } from './header.component';
import { SidebarComponent } from './sidebar.component';

@Component({
    selector: 'main-layout',
    imports: [HeaderComponent, SidebarComponent],
    templateUrl: './main.layout.html',
})
export class MainLayout {
}
