import React from 'react';
import PropTypes from 'prop-types';
import { Link } from '@reach/router';
import { Router } from '@reach/router';

import App from '@openmob/bluebird/src/components/layout/App';
import SideBar from '@openmob/bluebird/src/components/layout/sidebar/SideBar';
import ContentPanel from '@openmob/bluebird/src/components/layout/ContentPanel';
import SideBarHeader from '@openmob/bluebird/src/components/layout/sidebar/SideBarHeader';

import PageList from './PageList';
import PageEditor from './PageEditor';
import CauseEditor from './CauseEditor';
import PlaceHolder from './PlaceHolder';
import CauseMenu from './CauseMenu';
import FieldEditorMenu from './FieldEditorMenu';
import ImplementationEditorMenu from './ImplementationEditorMenu';
const CauseDashboard = () => (
  <App>
    {console.log('CauseDashboard ReRendering')}
    <SideBar>
      <SideBarHeader>
        <Link to="/">Open Mob</Link>
      </SideBarHeader>
      <Router>
        <FieldEditorMenu path="/pages/homepage/:pageID/component/:componentID/fields" />
        <CauseMenu path="/*" />
      </Router>
    </SideBar>
    <ContentPanel>
      <Router>
        <PlaceHolder path="/*" />
        <CauseEditor path="/" />
        <PageList path="/pages" />
        <PageEditor path="/pages/homepage/:pageID/*" />
      </Router>
    </ContentPanel>
  </App>
);

CauseDashboard.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]),
};

export default CauseDashboard;
