import { Component, Input } from '@angular/core';
import { FieldTree } from '@angular/forms/signals';

@Component({
    selector: 'app-form-input-error',
    imports: [],
    template: `
        @if (inputField().touched() && inputField().errors().length > 0) {
            <ul class="m-0 p-0 list-style-type-none">
                @for (error of inputField().errors(); track error) {
                    <li>
                        <small class="text-danger">{{ error.message }}</small>
                    </li>
                }
            </ul>
        }
    `,
})
export class FormInputError<T = unknown> {
    @Input() inputField!: FieldTree<T, string>
}
