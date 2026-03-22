import { HttpEvent, HttpHandlerFn, HttpRequest } from "@angular/common/http";
import { catchError, Observable, throwError } from 'rxjs';
import { environment } from '@environments/environment.development';


export function unauthorizedResponseMiddleware(
    req: HttpRequest<unknown>,
    next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
    return next(req).pipe(catchError(err => {
        if ([401].includes(err.status)) {
            window.location.href = environment.LOGIN_PAGE
        }

        const error = err.error?.message || err.statusText;
        return throwError(() => error);
    }))
}
