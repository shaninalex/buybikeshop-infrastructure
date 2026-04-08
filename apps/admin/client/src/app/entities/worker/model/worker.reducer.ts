import { createReducer, on } from '@ngrx/store';
import { WorkerModel } from './worker.model';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { actionWorkerSetList } from './worker.actions';

export interface WorkerState extends EntityState<WorkerModel> {
}

export const workersAdapter = createEntityAdapter<WorkerModel>();
export const initialState = workersAdapter.getInitialState();

export const workerReducer = createReducer(
    initialState,
    on(actionWorkerSetList, (state, action) => workersAdapter.addMany(action.workers, state)),
);

