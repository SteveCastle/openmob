import React from 'react';
import { Link } from '@reach/router';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';

import Menu from '@openmob/bluebird/src/components/menu/Menu';
import MenuItem from '@openmob/bluebird/src/components/menu/MenuItem';

function FieldEditorMenu({ causeID, pageID, componentID }) {
  return (
    <Menu vertical>
      <MenuItem>
        <Link to={`/app/cause/${causeID}/pages/homepage/${pageID}`}>
          <FontAwesomeIcon icon={faArrowLeft} />
          Cancel
        </Link>
      </MenuItem>
    </Menu>
  );
}

export default FieldEditorMenu;
