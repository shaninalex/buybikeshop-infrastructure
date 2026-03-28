import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MainLayout } from '@core';

@Component({
    selector: 'page-container',
    imports: [RouterOutlet, MainLayout],
    template: `
        <main-layout>
            <router-outlet />
        </main-layout>
    `,
})
export class PageContainer {}
