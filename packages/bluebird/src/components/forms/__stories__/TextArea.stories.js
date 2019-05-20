import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import TextArea from '../TextArea';

storiesOf('Forms/TextArea', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default TextArea', () => <TextArea onClick={action('clicked')} />);
