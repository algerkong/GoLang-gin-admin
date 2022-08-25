interface LoginData  {
    username: string
    password: string
}

interface UserData {
    ID?: number;
    CreatedAt?: string;
    UpdatedAt?: string;
    DeletedAt?: string;
    username: string;
    password?: string;
    age: number;
    role: number;
    avatar: string;
}

export type {
    LoginData,
    UserData
}