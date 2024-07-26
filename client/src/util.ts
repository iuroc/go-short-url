export const formHeader = { 'Content-Type': 'application/x-www-form-urlencoded' }

let user: User

/** 校验登录，可传入用户名和密码校验，也可以不传参而使用 Cookie 校验。
 * 
 * 校验成功后，将返回用户信息对象，校验失败则抛出错误。
 * 
 * @returns Error
 */
export const checkLogin = async (username?: string, password?: string) => {
    if (!username && !password && user) return user
    const res = await fetch('/api/user/login', {
        method: 'POST',
        headers: formHeader,
        body: username || password ? new URLSearchParams({
            username: username || '',
            password: password || ''
        }) : ''
    }).then(res => res.json()) as ResJSON<User>
    if (!res.success) throw new Error(res.message)
    user = res.data
    return res.data
}

export type ResJSON<Data = null> = {
    success: boolean
    data: Data
    message: string
}

export type User = {
    id: number
    username: string
    createTime: string
    updateTime: string
}