import React from 'react';
import { Link } from 'gatsby';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';

import Menu from '@openmob/bluebird/src/components/menu/Menu';
import MenuItem from '@openmob/bluebird/src/components/menu/MenuItem';

function ImplementationEditorMenu({ causeID, pageID, '*': currentPath }) {
  return (
    <Menu vertical>
      <MenuItem hide={currentPath === ''}>
        <Link to={`/app/cause/${causeID}/pages/homepage/${pageID}`}>
          <FontAwesomeIcon icon={faArrowLeft} />
          Cancel
        </Link>
      </MenuItem>
      <MenuItem hide={currentPath !== ''}>
        <Link to={`/app`}>
          <FontAwesomeIcon icon={faArrowLeft} />
          Implementation Editor
        </Link>
      </MenuItem>
    </Menu>
  );
}

export default ImplementationEditorMenu;
