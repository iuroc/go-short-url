import van from 'vanjs-core'
import LoginModel from './login'
import RegisterModel from './register'
import { Route } from 'vanjs-router'

const { div } = van.tags

export default () => {

    return Route({
        rule: /login|register/,
        Loader() {
            return div({ class: 'container' },
                div({ class: 'login-box px-3' },
                    LoginModel(), RegisterModel()
                )
            )
        },
    })
}
