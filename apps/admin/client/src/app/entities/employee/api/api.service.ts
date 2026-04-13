import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { EmployeeModel } from '@entities/employee';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { EmployeeCreateFormModel } from '@entities/employee/model/employee.model';

@Injectable({
    providedIn: 'root',
})
export class EmployeeApi {
    http = inject(HttpClient);

    GetEmployees(): Observable<EmployeeModel[]> {
        return this.http
            .get<APIResponse<EmployeeModel[]>>(`/api/v1/admin/employees`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    CreateEmployee(data: EmployeeCreateFormModel): Observable<EmployeeModel> {
        return this.http
            .post<APIResponse<EmployeeModel>>(`/api/v1/admin/employees/create`, data, {withCredentials: true})
            .pipe(map((response) => response.data));
    }
}
