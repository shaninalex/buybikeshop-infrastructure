import { Component, inject, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionProductGetList, ProductModel, selectProducts } from '@entities/product';
import { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-inventory-list',
    imports: [
        AsyncPipe
    ],
    templateUrl: './inventory-list.html',
})
export class InventoryList implements OnInit {
    private store = inject(Store);
    products$: Observable<ProductModel[]> = this.store.select(selectProducts);

    ngOnInit() {
        this.store.dispatch(actionProductGetList())
    }
}
