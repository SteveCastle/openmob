import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Button from '../Button';

storiesOf('Building Blocks/Buttons', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Basic Buttons', () => (
    <>
      <Button onClick={action('clicked')} label="Primary Button" />
      <Button
        onClick={action('clicked')}
        label="Outline Button"
        variant="outline"
      />
      <Button onClick={action('clicked')} label="Dark Button" variant="dark" />
      <Button
        onClick={action('clicked')}
        label="Warning Button"
        variant="dark"
      />
    </>
  ))
  .add('Button States', () => (
    <>
      <Button onClick={action('clicked')} label="Ready" />
      <Button onClick={action('clicked')} label="Loading" />
      <Button onClick={action('clicked')} label="Success" />
      <Button onClick={action('clicked')} label="Disabled" />
    </>
  ))
  .add('With confirmation modal', () => (
    <Button onClick={action('clicked')} label="Hello Button" />
  ))
  .add('With icons', () => (
    <Button onClick={action('clicked')} label="Hello Button" />
  ))
  .add('Block mode', () => (
    <Button onClick={action('clicked')} label="Hello Button" block />
  ));