import van from 'vanjs-core'
import { Route } from 'vanjs-router'
import Header from './header'

const { button, div, input } = van.tags

export default () => {
    return Route({
        rule: 'home',
        delayed: true,
        async onLoad() {
            // try {
            //     await checkLogin()
            //     this.show()
            // } catch {
            //     goto('login')
            // }
            this.show()
            setTimeout(() => {
                const input = this.element.childNodes[1].childNodes[0].childNodes[1].childNodes[0] as HTMLInputElement
                input.focus()
            })
        },
        Loader() {
            return div(
                Header(),
                div({ class: 'container px-4 px-sm-0' },
                    BigInput()
                )
            )
        },
    })
}

const BigInput = () => {
    return div({ class: 'big-input' },
        div({ class: 'mb-4 fs-2 text-center title fw-light' }, '请输入您需要缩短的网址'),
        div({ class: 'input-group' },
            input({ class: 'form-control form-control-lg' }),
            button({ class: 'btn btn-success btn-lg overflow-hidden px-md-5' }, '缩短')
        )
    )
}