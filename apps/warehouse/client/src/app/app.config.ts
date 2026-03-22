import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { ApplicationConfig, isDevMode, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideEffects } from '@ngrx/effects';
import { provideStore } from '@ngrx/store';
import { provideStoreDevtools } from '@ngrx/store-devtools';

import { effects, reducers, unauthorizedResponseMiddleware } from '@core';
import { routes } from '@pages';

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
