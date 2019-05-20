import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleBlogPost from '../SimpleBlogPost';

storiesOf('OpenMob Elements/SimpleBlogPost', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default SimpleBlogPost', () => (
    <SimpleBlogPost onClick={action('clicked')} />
  ));
