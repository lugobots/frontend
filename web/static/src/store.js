import { createStore, combineReducers, applyMiddleware } from 'redux'
import app from './redux/app/reducers'
import match from './redux/match/reducers'

const loggingMiddleware = (store) => (next) => (action) => {
  const previousState = store.getState().app.status
  next(action)
  console.log(`${previousState} -> ${store.getState().app.status }: dispatching %c${action.type}`, "color: blue" )
}

const rootReducer = combineReducers({
  app,
  match,
});

const store = createStore(rootReducer,
  applyMiddleware(
    loggingMiddleware,
  ));

export default store;