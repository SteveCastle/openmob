import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import TableHeaderCell from '../TableHeaderCell';

storiesOf('Building Blocks/TableHeaderCell', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default TableHeaderCell', () => (
      <TableHeaderCell onClick={action('clicked')} />
  ))
