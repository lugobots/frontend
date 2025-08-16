import { createStore, combineReducers, applyMiddleware } from 'redux'
import app from './redux/app/reducers'
import stadium from './redux/stadium/reducers'

const loggingMiddleware = (store) => (next) => (action) => {
  const previousAppState = store.getState().app.status
  const previousStadiumState = store.getState().stadium.status
  next(action)
  console.log(`App (${previousAppState} -> ${store.getState().app.status }) `+
          `Stadium (${previousStadiumState} -> ${store.getState().stadium.status }) `+
          `dispatching %c${action.type}`, "color: blue" )
}

const rootReducer = combineReducers({
  app,
  stadium,
});

const store = createStore(rootReducer,
  applyMiddleware(
    loggingMiddleware,
  ));

export default store;
