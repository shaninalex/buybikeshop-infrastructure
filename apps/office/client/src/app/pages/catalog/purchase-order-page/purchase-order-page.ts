import { Component } from '@angular/core';

@Component({
    selector: 'app-purchase-order-page',
    imports: [],
    template: `
        <div class="d-flex justify-content-between align-items-center mb-4">
            <div>
                <h2 class="mb-0">Add New Product / Variant</h2>
                <small class="text-muted">Purchase order for supplier</small>
            </div>
            <div>
                <button class="btn btn-outline-secondary me-2">Cancel</button>
                <button class="btn btn-primary">Save Product</button>
            </div>
        </div>

        <div class="row g-4">
            <div class="col-md-6">
                <div class="card">
                    <div class="card-header">Product Info</div>
                    <div class="card-body">
                        <div class="mb-3">
                            <label class="form-label">Product Name</label>
                            <input type="text" class="form-control" placeholder="Enter product name">
                        </div>
                        <div class="mb-3">
                            <label class="form-label">SKU / Code</label>
                            <input type="text" class="form-control" placeholder="Enter SKU or code">
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Category</label>
                            <select class="form-select">
                                <option selected>Electronics</option>
                                <option>Clothing</option>
                                <option>Accessories</option>
                            </select>
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Supplier</label>
                            <select class="form-select">
                                <option selected>Acme Corporation</option>
                                <option>Global Supplies</option>
                            </select>
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Unit Price</label>
                            <input type="number" class="form-control" placeholder="Enter price" step="0.01">
                        </div>
                    </div>
                </div>

                <div class="card mt-4">
                    <div class="card-header">Variant Options</div>
                    <div class="card-body">
                        <div class="mb-3">
                            <label class="form-label">Color</label>
                            <input type="text" class="form-control" placeholder="Enter color variant">
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Size</label>
                            <input type="text" class="form-control" placeholder="Enter size / dimension">
                        </div>
                        <div class="mb-3">
                            <label class="form-label">Additional SKU (optional)</label>
                            <input type="text" class="form-control" placeholder="Variant-specific SKU">
                        </div>
                    </div>
                </div>
            </div>

            <div class="col-md-6">
                <div class="card">
                    <div class="card-header">Purchase Order Quantities</div>
                    <div class="card-body">

                        <table class="table">
                            <thead class="table-light">
                            <tr>
                                <th>Variant</th>
                                <th>Unit</th>
                                <th>Quantity</th>
                                <th></th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr>
                                <td>Default</td>
                                <td>pcs</td>
                                <td>
                                    <input type="number" class="form-control form-control-sm" value="10">
                                </td>
                                <td>
                                    <button class="btn btn-sm btn-outline-danger">Remove</button>
                                </td>
                            </tr>
                            </tbody>
                        </table>

                        <div class="d-flex justify-content-end mt-2">
                            <button class="btn btn-sm btn-outline-primary">Add Variant Quantity</button>
                        </div>
                    </div>
                </div>

                <div class="card mt-4">
                    <div class="card-header">Notes</div>
                    <div class="card-body">
                        <textarea class="form-control" rows="4"
                                  placeholder="Add notes for supplier or internal team..."></textarea>
                        <div class="d-flex justify-content-end mt-2">
                            <button class="btn btn-primary btn-sm">Save Notes</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    `,
})
export class PurchaseOrderPage {
}
