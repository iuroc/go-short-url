import van from 'vanjs-core'
import Login from './routes/login/main'
import Home from './routes/home/main'

const { div } = van.tags

export default () => div(
    Login(), Home()
)
