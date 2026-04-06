import { createAction, props } from '@ngrx/store';
import { WorkerModel } from './worker.model';

export const actionWorkerGetList = createAction('[Worker] get list');
export const actionWorkerSetList = createAction(
    '[Worker] set list',
    props<{ workers: WorkerModel[] }>(),
);
