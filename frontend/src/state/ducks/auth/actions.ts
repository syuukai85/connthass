import { ActionTypes } from './types';
import firebase from 'firebase';
import { Action } from 'redux';

interface IsLoginAction extends Action {
  type: typeof ActionTypes.IS_LOGIN;
  payload: {
    uid: string;
    displayName: string | null;
    email: string | null;
  };
}

const isLogin = (user: firebase.User): IsLoginAction => {
  return {
    type: ActionTypes.IS_LOGIN,
    payload: {
      displayName: user.displayName,
      email: user.email,
      uid: user.uid
    }
  };
};

interface LogoutAction extends Action {
  type: typeof ActionTypes.LOGOUT;
}

const logout = (): LogoutAction => {
  return {
    type: ActionTypes.LOGOUT
  };
};

export default { isLogin, logout };
export type AuthAction = IsLoginAction | LogoutAction;
