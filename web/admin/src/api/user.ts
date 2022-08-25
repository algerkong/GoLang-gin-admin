import { type } from 'os'
import http from './http'
import { LoginData, UserData } from '@/types/user'
import { PageData } from '@/types/page'
import { AxiosAdapter, AxiosRequestConfig } from 'axios'
import { IDataType } from '@/types'


const Login = (data: LoginData) => {
    return http.post('/login', data)
}

const GetUsers = (data:(PageData|undefined)) => {
    return http.get<IDataType<UserData[]>>('/users',{
        params: data
    })
}

const AddUser = (data:UserData) => {
    return http.post<IDataType>('/user', data)
}

const UpdateUser = (id:number,data:UserData) => {
    return http.put<IDataType>('/user/'+id, data)
}

const DeleteUser = (id:number) => {
    return http.delete<IDataType>('/user/'+id)
}

const user = {
    Login,
    GetUsers,
    AddUser,
    UpdateUser,
    DeleteUser
}
export default user