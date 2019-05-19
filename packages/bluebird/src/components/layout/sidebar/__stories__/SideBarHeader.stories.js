import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SideBarHeader from '../SideBarHeader';

storiesOf('Layout System/SideBarHeader', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SideBarHeader', () => (
      <SideBarHeader onClick={action('clicked')} />
  ))
