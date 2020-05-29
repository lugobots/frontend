import { createStore } from 'redux'
import app from './redux/app/reducers'

const store = createStore(app)
export default store;