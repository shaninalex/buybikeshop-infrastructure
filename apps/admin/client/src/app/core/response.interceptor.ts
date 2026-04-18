import { HttpEvent, HttpHandlerFn, HttpRequest } from "@angular/common/http";
import { catchError, Observable, throwError } from 'rxjs';
import { environment } from '@environments/environment.development';
import { ApiError } from '@shared/models';

export function unauthorizedResponseMiddleware(
    req: HttpRequest<unknown>,
    next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
    return next(req).pipe(catchError(err => {
        if (err.status === 401) {
            window.location.href = environment.LOGIN_PAGE
        }

        const errors: ApiError[] = err.error?.errors ?? [{
            message: err.error?.message ?? err.statusText,
            reason: '',
            code: err.status,
            status: err.statusText,
        }];
        return throwError(() => errors);
    }))
}
