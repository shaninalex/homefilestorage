export interface User {
    id?: number
    first_name: string
    last_name: string
    email: string
    lang: string
    phone: string
    password?: string
}

export interface FileItem {
    id?: number
    name: string
    size: number
    modified: Date
    public: boolean
}