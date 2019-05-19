import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SubHeader from '../SubHeader';

storiesOf('Layout/SubHeader', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SubHeader', () => (
      <SubHeader onClick={action('clicked')} />
  ))
