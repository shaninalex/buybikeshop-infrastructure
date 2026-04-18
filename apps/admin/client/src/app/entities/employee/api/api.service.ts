import { inject, Injectable } from '@angular/core';
import { catchError, map, Observable, throwError } from 'rxjs';
import { EmployeeModel } from '@entities/employee';
import { ApiError, APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { EmployeeCreateFormModel } from '@entities/employee/model/employee.model';

@Injectable({
    providedIn: 'root',
})
export class EmployeeApi {
    http = inject(HttpClient);

    GetEmployees(): Observable<EmployeeModel[]> {
        return this.http
            .get<APIResponse<EmployeeModel[]>>(`/api/v1/admin/employees`, {withCredentials: true})
            .pipe(
                map((response) => response.data),
                catchError((errors: ApiError[]) => throwError(() => errors)),
            );
    }

    CreateEmployee(data: EmployeeCreateFormModel): Observable<EmployeeModel> {
        return this.http
            .post<APIResponse<EmployeeModel>>(`/api/v1/admin/employees/create`, data, {withCredentials: true})
            .pipe(
                map((response) => {
                    if (!response.status) throw response.errors;
                    return response.data;
                }),
                catchError((errors: ApiError[]) => throwError(() => errors)),
            );
    }

    UpdateEmployee(id: string, data: EmployeeCreateFormModel): Observable<EmployeeModel> {
        return this.http
            .patch<APIResponse<EmployeeModel>>(`/api/v1/admin/employees/${id}`, data, {withCredentials: true})
            .pipe(
                map((response) => {
                    if (!response.status) throw response.errors;
                    return response.data;
                }),
                catchError((errors: ApiError[]) => throwError(() => errors)),
            );
    }
}
