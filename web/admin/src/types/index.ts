export interface IDataType<T=any>{
    code: number;
    data?: T;
    msg: string;
    total: number | undefined;
}