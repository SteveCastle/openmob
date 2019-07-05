import React from 'react';
import { Router } from '@reach/router';
import ApolloClient from 'apollo-boost';
import { ApolloProvider } from 'react-apollo-hooks';
import Admin from '../../components/admin/Admin';
import Home from '../../components/admin/Home';
import CauseDashboard from '../../components/admin/CauseDashboard';
import ThemeProvider from '@openmob/bluebird/src/ThemeProvider';
import skyward from '@openmob/bluebird/src/themes/skyward';
import './reset.css';

const client = new ApolloClient({
  uri: 'http://localhost:4000',
});

function App() {
  return (
    <ThemeProvider theme={skyward}>
      <ApolloProvider client={client}>
        <Router>
          <Home path="app/*" />
          <CauseDashboard path="app/cause/:causeID/*" />
          <Admin path="app/admin/*" />
        </Router>
      </ApolloProvider>
    </ThemeProvider>
  );
}

export default App;
