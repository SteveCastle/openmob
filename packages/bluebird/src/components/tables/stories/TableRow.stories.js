import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import TableRow from '../TableRow';

storiesOf('Layout/TableRow', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default TableRow', () => (
      <TableRow onClick={action('clicked')} />
  ))
