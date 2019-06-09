import React from 'react';
import PropTypes from 'prop-types';
import { Link } from 'gatsby';
import { Router } from '@reach/router';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faArrowLeft,
  faFile,
  faPaperPlane,
  faHandsHelping,
  faMoneyCheckAlt,
  faCalendar,
  faPersonBooth,
  faBuilding,
} from '@fortawesome/free-solid-svg-icons';

import App from '@openmob/bluebird/src/components/layout/App';
import SideBar from '@openmob/bluebird/src/components/layout/sidebar/SideBar';
import ContentPanel from '@openmob/bluebird/src/components/layout/ContentPanel';
import SideBarHeader from '@openmob/bluebird/src/components/layout/sidebar/SideBarHeader';
import Menu from '@openmob/bluebird/src/components/menu/Menu';
import MenuItem from '@openmob/bluebird/src/components/menu/MenuItem';
import PageList from './PageList';
import PageEditor from './PageEditor';
import CauseEditor from './CauseEditor';
import PlaceHolder from './PlaceHolder';

const CauseDashboard = ({ causeID, '*': currentPath }) => (
  <App>
    <SideBar>
      <SideBarHeader>
        <Link to="/">Open Mob</Link>
      </SideBarHeader>
      <Menu vertical>
        <MenuItem hide={currentPath === ''}>
          <Link to={`/app/cause/${causeID}`}>
            <FontAwesomeIcon icon={faArrowLeft} />
            Back to Cause
          </Link>
        </MenuItem>
        <MenuItem hide={currentPath !== ''}>
          <Link to={`/app`}>
            <FontAwesomeIcon icon={faArrowLeft} />
            Back to Dashboard
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'pages'}>
          <Link to={`/app/cause/${causeID}/pages`}>
            <FontAwesomeIcon icon={faFile} />
            Pages
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'outreach'}>
          <Link to={`/app/cause/${causeID}/outreach`}>
            <FontAwesomeIcon icon={faPaperPlane} />
            Outreach
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'volunteering'}>
          <Link to={`/app/cause/${causeID}/volunteering`}>
            <FontAwesomeIcon icon={faHandsHelping} />
            Volunteer Coordination
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'fundraising'}>
          <Link to={`/app/cause/${causeID}/fundraising`}>
            <FontAwesomeIcon icon={faMoneyCheckAlt} />
            Fundraising
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'events'}>
          <Link to={`/app/cause/${causeID}/events`}>
            <FontAwesomeIcon icon={faCalendar} />
            Events
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'elections'}>
          <Link to={`/app/cause/${causeID}/elections`}>
            <FontAwesomeIcon icon={faPersonBooth} />
            Elections
          </Link>
        </MenuItem>
        <MenuItem active={currentPath === 'companies'}>
          <Link to={`/app/cause/${causeID}/companies`}>
            <FontAwesomeIcon icon={faBuilding} />
            Companies
          </Link>
        </MenuItem>
      </Menu>
    </SideBar>
    <ContentPanel>
      <Router>
        <CauseEditor path="/" />
        <PageList path="/pages" />
        <PageEditor path="/pages/homepage/:pageID" />
        <PlaceHolder path="/*" />
      </Router>
    </ContentPanel>
  </App>
);

CauseDashboard.propTypes = {
  children: PropTypes.node.isRequired,
};

export default CauseDashboard;
