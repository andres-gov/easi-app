// EXAMPLE - TO BE DELETED

import { UPDATE_DEMO_NAME } from 'constants/templateConstants';
import { DemoNameInterface } from 'types/demoName';

const initialState: DemoNameInterface = {
  name: ''
};

function templateReducer(state = initialState, action: any): DemoNameInterface {
  switch (action.type) {
    case UPDATE_DEMO_NAME:
      return {
        ...state,
        name: action.payload
      };
    default:
      return state;
  }
}

export default templateReducer;