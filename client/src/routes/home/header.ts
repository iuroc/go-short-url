import van from 'vanjs-core'

const { a, button, div } = van.tags

export default () => div({ class: 'border-bottom' },
    div({ class: 'container hstack p-4 px-sm-0' },
        div({ class: 'fs-4 fw-bold' }, 'Go Short URL 短网址平台'),
        div({ class: 'hstack gap-3 ms-auto' },
            a({ class: 'btn btn-primary', href: '#/login' }, '登录'),
            a({ class: 'btn btn-light border', href: '#/register' }, '注册')
        )
    )
)