import { createStore } from 'redux'
import game from './reducers'

const store = createStore(game)
export default store;