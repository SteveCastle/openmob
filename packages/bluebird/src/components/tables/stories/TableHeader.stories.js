import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import TableHeader from '../TableHeader';

storiesOf('Layout/TableHeader', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default TableHeader', () => (
      <TableHeader onClick={action('clicked')} />
  ))
