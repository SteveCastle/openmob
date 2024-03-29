import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'gatsby';
import { Router } from '@reach/router';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faPlus,
  faPlug,
  faCog,
  faArrowLeft,
  faOilCan,
} from '@fortawesome/free-solid-svg-icons';

import App from '@openmob/bluebird/src/components/layout/App';
import SideBar from '@openmob/bluebird/src/components/layout/sidebar/SideBar';
import ContentPanel from '@openmob/bluebird/src/components/layout/ContentPanel';
import SideBarHeader from '@openmob/bluebird/src/components/layout/sidebar/SideBarHeader';
import Menu from '@openmob/bluebird/src/components/menu/Menu';
import MenuItem from '@openmob/bluebird/src/components/menu/MenuItem';
import MyCauses from './MyCauses';
import New from './New';
import PlaceHolder from './PlaceHolder';
const Home = ({ '*': currentPath }) => (
  <App>
    <SideBar>
      <SideBarHeader>
        <Link to="/">Open Mob</Link>
      </SideBarHeader>
      <Menu vertical>
        <MenuItem hide={currentPath === ''}>
          <Link to="/app">
            <FontAwesomeIcon icon={faArrowLeft} />
            Back to Dashboard
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'new'}>
          <Link to="/app/new">
            <FontAwesomeIcon icon={faPlus} />
            Create Cause
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'plugins'}>
          <Link to="/app/plugins">
            <FontAwesomeIcon icon={faPlug} />
            Plugins
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'settings'}>
          <Link to="/app/settings">
            <FontAwesomeIcon icon={faCog} />
            Settings
          </Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/admin">
            <FontAwesomeIcon icon={faOilCan} />
            CRUD
          </Link>
        </MenuItem>
      </Menu>
    </SideBar>
    <ContentPanel>
      <Router>
        <MyCauses path="/*" />
        <New path="/new" />
        <PlaceHolder path="/plugins" />
        <PlaceHolder path="/settings" />
      </Router>
    </ContentPanel>
  </App>
);

Home.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]),
};

export default Home;
