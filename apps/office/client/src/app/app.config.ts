import { ApplicationConfig, isDevMode, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { provideRouter } from '@angular/router';

import { provideStore } from '@ngrx/store';
import { provideStoreDevtools } from '@ngrx/store-devtools';
import { provideEffects } from '@ngrx/effects';
import { routes } from '@pages';
import { effects, reducers, unauthorizedResponseMiddleware } from '@core';

export const appConfig: ApplicationConfig = {
    providers: [
        provideBrowserGlobalErrorListeners(),
        provideHttpClient(withInterceptors([unauthorizedResponseMiddleware])),
        provideRouter(routes),
        provideStore(reducers),
        provideEffects(effects),
        provideStoreDevtools({ maxAge: 25, logOnly: !isDevMode() }),
    ],
};
