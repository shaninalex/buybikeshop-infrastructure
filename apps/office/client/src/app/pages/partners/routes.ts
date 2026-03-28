import { Routes } from "@angular/router";
import { PartnersContainer } from "@pages/partners/partners.container";
import { PartnersListPage } from "@pages/partners/partners-list-page";
import { PartnerDetailPage } from "@pages/partners/partner-detail-page";
import { PartnersCreatePage } from "@pages/partners/partners-create-page";

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
                path: ":id",
                component: PartnerDetailPage,
            }
        ]
    }
];
