import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SideBar from '../SideBar';

storiesOf('Layout/SideBar', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SideBar', () => (
      <SideBar onClick={action('clicked')} />
  ))
