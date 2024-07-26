import van from 'vanjs-core'
import { Route, goto } from 'vanjs-router'
import { checkLogin } from '../../util'
import { checkPasswordFormat, checkUsernameFormat } from './util'
import { FooterLink, GoButton, Input, Title } from './view'

const { div } = van.tags

export default () => {
    const username = van.state('')
    const password = van.state('')
    const usernameInvalid = van.state('')
    const passwordInvalid = van.state('')
    const clickLogin = async () => {
        try {
            if (username.val == '')
                throw new Error('用户名不能为空')
            if (username.val.length > 50)
                throw new Error('用户名过长')
        } catch (error) {
            if (error instanceof Error) {
                return usernameInvalid.val = error.message
            }
        }
        try {
            if (password.val == '')
                throw new Error('用户名不能为空')
            if (password.val.length > 50)
                throw new Error('密码过长')
        } catch (error) {
            if (error instanceof Error) {
                return passwordInvalid.val = error.message
            }
        }
        try {
            await checkLogin(username.val, password.val)
            goto('home')
        } catch (error) {
            if (error instanceof Error) alert(error.message)
        }
    }
    return Route<HTMLDivElement>({
        rule: 'login',
        Loader() {
            return div({ class: 'vstack gap-4' },
                Title('用户登录', 'success'),
                Input(username, '请输入用户名', 'text', usernameInvalid),
                Input(password, '请输入密码', 'password', passwordInvalid),
                GoButton('登录', 'success', clickLogin),
                FooterLink('success', '#/register', '没有账号？点此注册')
            )
        },
    })
}
