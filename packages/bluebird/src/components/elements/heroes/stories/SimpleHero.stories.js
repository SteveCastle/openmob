import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleHero from '../SimpleHero';

storiesOf('Layout/SimpleHero', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SimpleHero', () => (
      <SimpleHero onClick={action('clicked')} />
  ))
