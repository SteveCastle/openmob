import React from 'react';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faDog } from '@fortawesome/free-solid-svg-icons';

const PlaceHolder = () => (
  <div
    style={{
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      height: '100%',
    }}
  >
    <FontAwesomeIcon icon={faDog} size="6x" spin />
  </div>
);

export default PlaceHolder;
