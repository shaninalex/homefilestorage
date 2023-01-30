export interface LoginResponse {
    access: string
    refresh: string
}

export interface LoginRequest {
    email: string
    password: string
}


export interface RegisterRequest {
    email: string
    password: string
    password_confirm: string
}
