/** 校验用户名格式
 * @throws Error
 */
export const checkUsernameFormat = (username: string) => {
    if (username.match(/^\w{3,20}$/) == null)
        throw new Error('用户名格式错误，要求 3-20 位，可使用数字、字母、下划线')
}

/** 校验密码格式
 * @throws Error
 */
export const checkPasswordFormat = (password: string) => {
    if (password.match(/^[\x00-\x7F]{8,20}$/) == null)
        throw new Error('密码格式错误，要求 8-20 位，可使用数字、字母、特殊符号')
    return true
}