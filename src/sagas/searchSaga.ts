import axios from 'axios';
import { takeLatest, call, put } from 'redux-saga/effects';
import { GET_ALL_SYSTEM_SHORTS } from '../constants/system';
import { putSystemShorts } from '../actions/searchActions';

function requestSystemShorts(accessToken: string) {
  return axios.get('http://localhost:8080/api/v1/systems', {
    headers: {
      Authorization: `Bearer ${accessToken}`
    }
  });
}

export function* fetchAllSystemShorts(action: any) {
  const obj = yield call(requestSystemShorts, action.accessToken);
  yield put(putSystemShorts(obj.data));
}

export function* searchSaga() {
  yield takeLatest(GET_ALL_SYSTEM_SHORTS, fetchAllSystemShorts);
}