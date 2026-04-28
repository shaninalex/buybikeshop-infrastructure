export interface APIResponse<T> {
    messages: string[];
    status: boolean;
    data: T;
    errors: ApiError[];
}

export interface ApiError {
    message: string
    reason: string
    code: number
    status: string
}
