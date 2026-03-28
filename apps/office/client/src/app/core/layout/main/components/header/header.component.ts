import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';

@Component({
    selector: 'app-header',
    imports: [
        RouterLink,
        RouterLinkActive
    ],
    template: `
        <div class="mb-4" style="height: 56px;"></div>
        <nav class="navbar navbar-expand-lg bg-dark position-fixed top-0 left-0 w-100" data-bs-theme="dark">
            <div class="container">
                <a class="navbar-brand" routerLink="/">
                    <img src="images/logo.png" alt="BuyBikeShop Office" style="width: 2rem">
                </a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                        data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
                        aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item">
                            <a class="nav-link" routerLinkActive="active fw-bold"
                               [routerLinkActiveOptions]="{ exact: true }"
                               aria-current="page" routerLink="/">Home</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" routerLinkActive="active fw-bold" routerLink="/partners">Partners</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" routerLinkActive="active fw-bold" routerLink="/catalog">Catalog</a>
                        </li>
                    </ul>
                    <ul class="navbar-nav">
                        <li class="nav-item dropdown">
                            <span class="nav-link" data-bs-toggle="dropdown">
                                <i class="fa-solid fa-bell"></i>
                                <span class="badge text-bg-danger">4</span>
                            </span>
                        </li>
                        <li class="nav-item dropdown">
                            <span class="nav-link py-0">
                                <img src="images/default-avatar.png" alt="profile" style="width: 2rem">
                            </span>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    `
})
export class HeaderComponent {
}
