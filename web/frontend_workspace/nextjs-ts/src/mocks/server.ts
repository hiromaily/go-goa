//import { setupServer } from 'msw/node'
import { handlers } from './handlers'

//export const server = setupServer(...handlers)
export const server = require('msw/node').setupServer(...handlers)
