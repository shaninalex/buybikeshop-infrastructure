import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
    selector: 'app-employee-container',
    imports: [RouterOutlet],
    template: `<router-outlet />`,
})
export class EmployeeContainer {}
