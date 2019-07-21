import { takeEvery, call, put } from 'redux-saga/effects';
import actions from './actions';
import { EventApi, SearchEventsRequest } from 'typescript-fetch-api';
import { ActionTypes } from './types';

let api = new EventApi();

function* searchRecentlyAddedEvent() {
  const searchEvents = (req: SearchEventsRequest) => {
    return api
      .searchEvents(req)
      .then(events => events)
      .catch(error => error);
  };
  try {
    const events = yield call(searchEvents, {});
    yield put(actions.searchSuccessRecentlyAddedEvent(events));
  } catch (error) {
    yield put(actions.searchErrorRecentlyAddedEvent(error));
  }
}

const sagas = [takeEvery(ActionTypes.REQUEST_NEWLY_EVENT, searchRecentlyAddedEvent)];
export default sagas;
