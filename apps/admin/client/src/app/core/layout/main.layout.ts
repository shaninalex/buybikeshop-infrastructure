import { Component } from '@angular/core';
import { HeaderComponent } from './components/header.component';
import { SidebarComponent } from './components/sidebar.component';

@Component({
    selector: 'main-layout',
    imports: [HeaderComponent, SidebarComponent],
    templateUrl: './main.layout.html',
    styleUrl: './main.layout.css',
})
export class MainLayout {
}
