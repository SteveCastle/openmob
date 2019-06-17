import React from 'react';
import { Link } from 'gatsby';
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

import Menu from '@openmob/bluebird/src/components/menu/Menu';
import MenuItem from '@openmob/bluebird/src/components/menu/MenuItem';

function CauseMenu({ causeID, '*': currentPath }) {
  return (
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
  );
}

export default CauseMenu;
