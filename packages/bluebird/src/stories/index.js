import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import { linkTo } from '@storybook/addon-links';

import Button from '../components/buttons/Button';

storiesOf('Button', module)
  .add('Defaults', () => (
    <Button onClick={action('clicked')} label="Hello Button" />
  ))
  .add('Block mode', () => (
    <Button onClick={action('clicked')} label="Hello Button" block />
  ))
  .add('Dark Mode', () => (
    <Button onClick={action('clicked')} label="Hello Button" dark />
  ));
