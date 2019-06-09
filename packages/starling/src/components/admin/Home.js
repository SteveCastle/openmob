import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'gatsby';
import { Router } from '@reach/router';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faPlus,
  faPlug,
  faCog,
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

const Home = () => (
  <App>
    <SideBar>
      <SideBarHeader>
        <Link to="/">Open Mob</Link>
      </SideBarHeader>
      <Menu vertical>
        <MenuItem>
          <Link to="/app/new">
            <FontAwesomeIcon icon={faPlus} />
            Create Cause
          </Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/plugins">
            <FontAwesomeIcon icon={faPlug} />
            Plugins
          </Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/settings">
            {' '}
            <FontAwesomeIcon icon={faCog} />
            Settings
          </Link>
        </MenuItem>
        <MenuItem>
          <Link to="/app/admin">
            {' '}
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
      </Router>
    </ContentPanel>
  </App>
);

Home.propTypes = {
  children: PropTypes.node.isRequired,
};

export default Home;
