import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import ImageGrid from '../ImageGrid';

storiesOf('Building Blocks/ImageGrid', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default ImageGrid', () => (
      <ImageGrid onClick={action('clicked')} />
  ))
