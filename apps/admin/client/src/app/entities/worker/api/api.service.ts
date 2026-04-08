import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { WorkerModel } from '@entities/worker';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

interface workersResponse {
    workers: WorkerModel[]
}

@Injectable({
    providedIn: 'root',
})
export class WorkerApi {
    http = inject(HttpClient);

    GetWorkers(): Observable<WorkerModel[]> {
        return this.http
            .get<APIResponse<workersResponse>>(`/api/v1/admin/workers`)
            .pipe(
                map((response) => response.data),
                map((response) => response.workers),
            );
    }
}
