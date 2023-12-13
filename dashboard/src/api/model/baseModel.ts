export interface BasicPageParams {
    page: number;
    pageSize: number;
}

export interface BasicFetchResult<T> {
    items: T[];
    total: number;
}

export interface Response {
    code: number;
    data: BasicFetchResult<any>;
    message: string;
}