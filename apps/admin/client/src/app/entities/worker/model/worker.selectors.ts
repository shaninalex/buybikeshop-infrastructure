import { createFeatureSelector, createSelector } from '@ngrx/store';
import { workersAdapter, WorkerState } from './worker.reducer';

export const selectWorkersFeature = createFeatureSelector<WorkerState>('worker');
export const workersSelectors = workersAdapter.getSelectors();

export const selectWorkers = createSelector(selectWorkersFeature, (state) =>
    workersSelectors.selectAll(state),
);

export const selectWorkerByID = (id: number) =>
    createSelector(selectWorkersFeature, (state: WorkerState) =>
        workersSelectors.selectAll(state).find((p) => p.id === id),
    );
