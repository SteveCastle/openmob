import { action } from '@storybook/addon-actions';
import { storiesOf } from '@storybook/react';
import React from 'react';
import Form from '../Form';
import Widget from '../Widget'
import Input from '../Input'
import Label from '../Label'
import TextArea from '../TextArea'

storiesOf('Forms/Form', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Form', () => (
    <Form onChange={action('change')} >
      <Widget>
        <Label>Form Item</Label>
        <Input onChange={action('change')} dark />
      </Widget>
      <Widget>
        <Label>Form Item</Label>
        <TextArea onChange={action('change')} dark />
      </Widget>
    </Form>))