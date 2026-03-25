import { Component } from '@angular/core';
import { RouterOutlet } from "@angular/router";

@Component({
    selector: "catalog-container",
    imports: [
        RouterOutlet
    ],
    template: `
        <router-outlet/>
    `
})
export class CatalogContainer {
}

