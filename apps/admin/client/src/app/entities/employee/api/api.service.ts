import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { EmployeeModel } from '@entities/employee';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

interface employeesResponse {
    employees: EmployeeModel[]
}

@Injectable({
    providedIn: 'root',
})
export class EmployeeApi {
    http = inject(HttpClient);

    GetEmployees(): Observable<EmployeeModel[]> {
        return this.http
            .get<APIResponse<employeesResponse>>(`/api/v1/admin/employees`)
            .pipe(
                map((response) => response.data),
                map((response) => response.employees),
            );
    }
}
