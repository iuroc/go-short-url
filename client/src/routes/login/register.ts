import { Route } from 'vanjs-router'
import { FooterLink, GoButton, Input, Title } from './view'
import van from 'vanjs-core'

const { div } = van.tags

export default () => {
    const username = van.state('')
    const password = van.state('')
    const password2 = van.state('')
    const usernameInvalid = van.state('')
    const passwordInvalid = van.state('')
    const password2Invalid = van.state('')
    const clickRegister = () => {

    }

    return Route<HTMLDivElement>({
        rule: 'register',
        Loader() {
            return div({ class: 'vstack gap-4' },
                Title('用户注册', 'primary'),
                Input(username, '请输入用户名', 'text', usernameInvalid),
                Input(password, '请输入密码', 'password', passwordInvalid),
                Input(password2, '请重复输入密码', 'password', password2Invalid),
                GoButton('注册', 'primary', clickRegister),
                FooterLink('primary', '#/login', '已有账号？点此登录')
            )
        },
    })
}
