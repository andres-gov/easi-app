import React from 'react';
import Header from 'components/Header';
import { useOktaAuth } from '@okta/okta-react';
import ActionBanner from '../../components/shared/ActionBanner/index';

// This view can be deleted whenever we're ready
// This is just a sandbox page for us to test things out

const onButtonClick = (authState: any, authService: any) => {
  const tokenManager = authService.getTokenManager();
  tokenManager.get('idToken').then((token: any) => {
    console.log('token', token)
  })
  console.log('authState', authState);
  alert(`I WAS CLICKED!`);
  return null
}

const Sandbox = () => {
  console.log(useOktaAuth);

  const { authState, authService }: { authState: object, authService: any } = useOktaAuth();

  return (
    <div>
      <Header />
      <div className="grid-container">
        <h1>Sandbox</h1>
          <ActionBanner
            title="thing"
            helpfulText="lots of helpful text"
            label="I am a button"
            onClick={() => onButtonClick(authState, authService)}
          />
      </div>
    </div>
  );
};

export default Sandbox;
