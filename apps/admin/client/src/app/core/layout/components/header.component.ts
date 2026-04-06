import { Component, inject } from '@angular/core';
import { ThemeSwitcher } from '@core/layout/components';
import { UiService } from '@shared/ui';
import { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';
import { environment } from '@environments/environment.development';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';

@Component({
    selector: 'app-header',
    template: `
        <nav class="navbar navbar-expand-lg bg-body-tertiary border-bottom">
            <div class="container-fluid">
                @if (title$ | async; as title) {
                    <a class="navbar-brand fw-bold" href="#">{{ title }}</a>
                }
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                        data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
                        aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav mb-2 mb-lg-0 ms-auto">
                        <li class="nav-item d-flex align-items-center">
                            <app-theme-switcher/>
                        </li>
                        <li class="nav-item dropdown">
                            <button [cdkMenuTriggerFor]="menu" class="nav-link">
                                <div class="d-flex align-items-center gap-2">
                                    profile
                                </div>
                            </button>

                            <ng-template #menu>
                                <div class="dropdown-menu d-block" cdkMenu style="right: -55px">
                                    <button (click)="logout()" class="dropdown-item" cdkMenuItem>Sign out</button>
                                </div>
                            </ng-template>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    `,
    host: {
        class: "header"
    },
    imports: [
        CdkMenuTrigger, CdkMenu, CdkMenuItem,
        ThemeSwitcher,
        AsyncPipe
    ]
})
export class HeaderComponent {
    private ui = inject(UiService)
    title$: Observable<string> = this.ui.pageTitle;

    logout(): void {
        // TODO: get logout token
        window.location.href = `${environment.AUTH_SERVER}/self-service/logout?token=<TODO: get logout token. Frontend can't directly get this>`
    }
}
