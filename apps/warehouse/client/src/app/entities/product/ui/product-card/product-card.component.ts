import { Component, Input } from '@angular/core';
import { ProductModel } from '@entities/product/model';

@Component({
    selector: 'app-ui-product-card',
    template: ` <div class="card"></div> `,
})
export class UiProductCard {
    @Input() product: ProductModel;
}
