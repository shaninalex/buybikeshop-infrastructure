import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
    selector: "partners-page",
    imports: [
        RouterLink
    ],
    template: `
        <h3>Partners page</h3>
        <nav class="navbar bg-body-tertiary mb-4">
            <div class="container-fluid">
                <form class="d-flex" role="search">
                    <input class="form-control me-2" type="search" placeholder="Search partner" aria-label="Search"/>
                    <button class="btn btn-outline-success" type="submit">Search</button>
                </form>
            </div>
        </nav>

        <table class="table">
            <thead>
            <tr>
                <th scope="col">#</th>
                <th scope="col">Name</th>
                <th scope="col"></th>
            </tr>
            </thead>
            <tbody>
            <tr>
                <th scope="row">1</th>
                <td><a routerLink="/partners/1" class="text-body">PartsParts</a></td>
                <td><a href="#" class="btn btn-sm btn-info">view</a></td>
            </tr>
            <tr>
                <th scope="row">2</th>
                <td><a routerLink="/partners/2" class="text-body">Shimano ltd.</a></td>
                <td><a href="#" class="btn btn-sm btn-info">view</a></td>
            </tr>
            <tr>
                <th scope="row">3</th>
                <td><a routerLink="/partners/3" class="text-body">John Doe</a></td>
                <td><a href="#" class="btn btn-sm btn-info">view</a></td>
            </tr>
            </tbody>
        </table>
    `
})
export class PartnersPage {
}

