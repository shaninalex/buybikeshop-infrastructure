import { Routes } from "@angular/router";
import { PartnersContainer } from "@pages/partners/partners.container";
import { PartnersListPage } from "@pages/partners/partners-list-page";
import { PartnerDetailPage } from "@pages/partners/partner-detail-page";
import { PartnersCreatePage } from "@pages/partners/partners-create-page";
import { PartnersRolesPage } from '@pages/partners/partners-roles-page';

export const partners_routes: Routes = [
    {
        path: "partners",
        component: PartnersContainer,
        children: [
            {
                path: "",
                component: PartnersListPage,
            },
            {
                path: "create",
                component: PartnersCreatePage,
            },
            {
                path: "roles",
                component: PartnersRolesPage,
            },
            {
                path: ":id",
                component: PartnerDetailPage,
            }
        ]
    }
];
