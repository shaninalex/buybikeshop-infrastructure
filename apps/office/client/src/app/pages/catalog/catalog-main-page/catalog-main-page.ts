import { Component } from '@angular/core';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'app-catalog-main-page',
    imports: [
        RouterLink
    ],
    template: `
        <div class="d-flex align-items-start justify-content-between">
            <h2>Catalog</h2>
            <a routerLink="/catalog/purchase-order" class="btn btn-primary">Purchase Order</a>
        </div>
    `,
})
export class CatalogMainPage {
}
