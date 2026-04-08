import { WorkerEffects, workerReducer } from "@entities/worker";

export const effects = [WorkerEffects];
export const reducers = {
    workers: workerReducer,
};
