import { Component, forwardRef, inject, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionPartnerRoleGetList, selectPartnerRoles } from '@entities/partner-role';
import { AsyncPipe } from '@angular/common';
import { NG_VALUE_ACCESSOR } from '@angular/forms';

@Component({
    selector: 'app-partner-role-selector',
    imports: [
        AsyncPipe
    ],
    providers: [
        {
            provide: NG_VALUE_ACCESSOR,
            useExisting: forwardRef(() => PartnerRoleSelector),
            multi: true
        }
    ],
    template: `
        @if (roles$ | async; as roles) {
            @for (role of roles; track role.id) {
                <div class="form-check">
                    <input
                        class="form-check-input"
                        type="checkbox"
                        [id]="'role-' + role.id"
                        [checked]="selectedIds.has(role.id)"
                        (change)="onToggle(role.id, $any($event.target).checked)" />

                    <label class="form-check-label" [for]="'role-' + role.id">
                        {{ role.role }}
                    </label>
                </div>
            }
        }
    `,
})
export class PartnerRoleSelector implements OnInit {
    private store = inject(Store);
    roles$ = this.store.select(selectPartnerRoles)

    ngOnInit() {
        this.store.dispatch(actionPartnerRoleGetList())
    }

    // internal state
    selectedIds = new Set<number>();

    // callbacks
    private onChange = (value: number[]) => {};
    private onTouched = () => {};

    // called when Angular sets initial value
    writeValue(value: number[] | null): void {
        this.selectedIds = new Set(value || []);
    }

    registerOnChange(fn: (value: number[]) => void): void {
        this.onChange = fn;
    }

    registerOnTouched(fn: () => void): void {
        this.onTouched = fn;
    }

    onToggle(id: number, checked: boolean) {
        if (checked) {
            this.selectedIds.add(id);
        } else {
            this.selectedIds.delete(id);
        }

        this.onChange(Array.from(this.selectedIds));
        this.onTouched();
    }

    protected readonly HTMLInputElement = HTMLInputElement;
}
