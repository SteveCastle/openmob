import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Row from '../Row';

storiesOf('Layout System/Row', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default Row', () => <Row onClick={action('clicked')} />);
