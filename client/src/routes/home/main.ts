import van from 'vanjs-core'
import { Route, goto } from 'vanjs-router'
import { checkLogin } from '../../util'

const { div } = van.tags

export default () => {
    return Route({
        rule: 'home',
        delayed: true,
        async onLoad() {
            try {
                await checkLogin()
                this.show()
            } catch {
                goto('login')
            }
        },
        Loader() {
            return div('Hello World')
        },
    })
}