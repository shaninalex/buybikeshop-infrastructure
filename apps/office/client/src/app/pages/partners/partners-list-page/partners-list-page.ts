import { Component } from '@angular/core';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'app-partners-list-page',
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

                <a routerLink="/partners/create" class="btn btn-primary">Create</a>
            </div>
        </nav>

        <table class="table">
            <thead>
            <tr>
                <th scope="col" style="width: 20px">#</th>
                <th scope="col">Name</th>
            </tr>
            </thead>
            <tbody>
            <tr>
                <th scope="row">1</th>
                <td><a routerLink="/partners/1" class="text-body">Acme Corporation</a></td>
            </tr>
            <tr>
                <th scope="row">2</th>
                <td><a routerLink="/partners/2" class="text-body">Shimano ltd.</a></td>
            </tr>
            <tr>
                <th scope="row">3</th>
                <td><a routerLink="/partners/3" class="text-body">John Doe</a></td>
            </tr>
            </tbody>
        </table>
    `,
})
export class PartnersListPage {
}
