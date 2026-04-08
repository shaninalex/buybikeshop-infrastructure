import { Actions, createEffect, ofType } from '@ngrx/effects';
import { inject, Injectable } from '@angular/core';
import { actionWorkerGetList, actionWorkerSetList, } from './worker.actions';
import { exhaustMap, of, switchMap } from 'rxjs';
import { WorkerApi } from '../api/api.service';

@Injectable()
export class WorkerEffects {
    private actions$ = inject(Actions);
    private workersApi = inject(WorkerApi);

    get_workers_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionWorkerGetList),
            exhaustMap(() =>
                this.workersApi
                    .GetWorkers()
                    .pipe(switchMap((data) => of(actionWorkerSetList({ workers: data })))),
            ),
        ),
    );
}
