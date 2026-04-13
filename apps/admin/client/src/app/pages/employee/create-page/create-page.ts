import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterLink } from '@angular/router';
import { UiService } from '@shared/ui';
import { GoogleForm, ManualForm } from '@pages/employee/create-page/components';

@Component({
    selector: 'app-employee-create-page',
    standalone: true,
    imports: [CommonModule, ReactiveFormsModule, RouterLink, ManualForm, GoogleForm],
    templateUrl: './create-page.html',
})
export class CreatePage implements OnInit {
    private fb = inject(FormBuilder)
    private ui = inject(UiService)

    ngOnInit() {
        this.ui.setPageTitle('Create Employee')
    }

    form = this.fb.group({
        creationMethod: ['manual', Validators.required],
    });

    get creationMethod(): 'manual' | 'google' {
        return this.form.get('creationMethod')?.value as 'manual' | 'google';
    }
}
