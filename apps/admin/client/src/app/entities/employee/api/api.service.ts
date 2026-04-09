import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { EmployeeModel } from '@entities/employee';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class EmployeeApi {
    http = inject(HttpClient);

    GetEmployees(): Observable<EmployeeModel[]> {
        return this.http
            .get<APIResponse<EmployeeModel[]>>(`/api/v1/admin/employees`)
            .pipe(map((response) => response.data));
    }
}
