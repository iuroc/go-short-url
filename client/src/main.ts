/// <reference types="vite/client" />
import van from 'vanjs-core'
import App from './app'
import 'bootstrap/dist/css/bootstrap.css'
import '../css/main.css'

van.add(document.body, App())