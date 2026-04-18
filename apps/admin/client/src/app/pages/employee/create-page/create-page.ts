import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterLink } from '@angular/router';
import { UiService } from '@shared/ui';
import { GoogleForm, ManualForm } from '@pages/employee/create-page/components';

@Component({
    selector: 'app-employee-create-page',
    standalone: true,
    imports: [CommonModule, ReactiveFormsModule, RouterLink, ManualForm, GoogleForm, FormsModule],
    templateUrl: './create-page.html',
})
export class CreatePage implements OnInit {
    private ui = inject(UiService)
    creationMethod: 'manual' | 'google' = 'manual'

    ngOnInit() {
        this.ui.setPageTitle('Create Employee')
    }
}
