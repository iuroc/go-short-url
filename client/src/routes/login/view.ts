import van, { State } from 'vanjs-core'

const { a, button, div, input } = van.tags

export const FooterLink = (color: string, href: string, text: string) => div({ class: 'text-center' }, a({
    class: `text-decoration-none focus-ring py-1 px-2 rounded link-${color}`,
    href: href,
    style: `--bs-focus-ring-color: var(--bs-${color}-border-subtle);`
}, text))

export const Input = (
    value: State<string>,
    placeholder: string,
    type: 'text' | 'password',
    invalid: State<string>,
) => div(
    input({
        class: () => `form-control border-2 ${invalid.val ? 'is-invalid' : ''}`, value, placeholder, type,
        ...(type == 'password' ? { autocomplete: 'new-password' } : {}),
        oninput: event => {
            invalid.val = ''
            return value.val = event.target.value
        }
    }),
    div({ class: 'invalid-feedback' }, invalid)
)

export const Title = (text: string, color: string) => div({ class: `fs-1 text-center text-${color} title` }, text)

export const GoButton = (text: string, color: string, onclick: () => void) => div({ class: 'd-grid' }, button({ class: `btn btn-${color}`, onclick }, text))