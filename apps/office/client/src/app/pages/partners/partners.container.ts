import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from "@angular/router";

@Component({
    selector: "partners-container",
    imports: [
        RouterOutlet,
        RouterLink,
        RouterLinkActive
    ],
    template: `
        <nav class="navbar navbar-expand-lg bg-body-tertiary mb-4">
            <div class="container-fluid">
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <a routerLink="/partners" routerLinkActive="active" [routerLinkActiveOptions]="{ exact: true }" class="nav-link">Partners</a>
                    </li>
                    <li class="nav-item">
                        <a routerLink="/partners/create" routerLinkActive="active" [routerLinkActiveOptions]="{ exact: true }" class="nav-link">Create Partner</a>
                    </li>
                    <li class="nav-item">
                        <a routerLink="/partners/roles" routerLinkActive="active" [routerLinkActiveOptions]="{ exact: true }" class="nav-link" href="#">Roles</a>
                    </li>
                </ul>
            </div>
        </nav>

        <router-outlet/>
    `
})
export class PartnersContainer {
}

