import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';

import Button from '../components/buttons/Button';
import Input from '../components/forms/Input';
import Header from '../components/type/Header';
import Paragraph from '../components/type/Paragraph';

import Layout from '../components/layout/Layout';
import Row from '../components/layout/Row';
import Column from '../components/layout/Column';
import Container from '../components/layout/Container';

storiesOf('Typography', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Header', () => <Header>Nice Header</Header>)
  .add('Paragraph', () => <Paragraph>A nice paragraph.</Paragraph>);

storiesOf('Layout', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Layout', () => (
    <Layout tracing={5}>
      <Row tracing={5}>
        <Column tracing={5} />
        <Column tracing={5} />
      </Row>
    </Layout>
  ))
  .add('Container', () => (
    <Layout tracing={5}>
      <Container tracing={5}>
        <Row tracing={5}>
          <Column tracing={5} />
        </Row>
      </Container>
    </Layout>
  ));

storiesOf('Buttons', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Defaults', () => (
    <Button onClick={action('clicked')} label="Hello Button" />
  ))
  .add('Block mode', () => (
    <Button onClick={action('clicked')} label="Hello Button" block />
  ))
  .add('Dark Mode', () => (
    <Button onClick={action('clicked')} label="Hello Button" dark />
  ));

storiesOf('Forms', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Form', () => <Input onClick={action('clicked')} />)
  .add('Input', () => <Input onClick={action('clicked')} block />)
  .add('DropDown', () => <Input onClick={action('clicked')} dark />);
