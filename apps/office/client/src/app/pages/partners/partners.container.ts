import { Component } from '@angular/core';
import { RouterOutlet } from "@angular/router";

@Component({
    selector: "partners-container",
    imports: [
        RouterOutlet
    ],
    template: `
        <router-outlet/>
    `
})
export class PartnersContainer {
}

