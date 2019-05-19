import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import TableCell from '../TableCell';

storiesOf('Building Blocks/TableCell', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default TableCell', () => (
      <TableCell onClick={action('clicked')} />
  ))
