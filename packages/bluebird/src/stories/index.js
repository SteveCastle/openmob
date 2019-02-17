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
import Form from '../components/forms/Form';
import Masonry from '../components/photos/Masonry';

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
        <Column tracing={5} size={12} />
      </Row>
      <Row tracing={5}>
        <Column tracing={5} size={6} />
        <Column tracing={5} size={6} />
      </Row>
      <Row tracing={5}>
        <Column tracing={5} size={6} />
        <Column tracing={5} size={6} />
        <Column tracing={5} size={6} />
      </Row>
      <Row tracing={5}>
        <Column tracing={5} size={3} />
        <Column tracing={5} size={3} />
        <Column tracing={5} size={3} />
        <Column tracing={5} size={3} />
      </Row>
      <Row tracing={5}>
        <Column tracing={5} size={2} />
        <Column tracing={5} size={2} />
        <Column tracing={5} size={2} />
        <Column tracing={5} size={2} />
        <Column tracing={5} size={2} />
        <Column tracing={5} size={2} />
      </Row>
      <Row tracing={5}>
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
        <Column tracing={5} size={1} />
      </Row>
    </Layout>
  ))
  .add('Container', () => (
    <Layout tracing={5}>
      <Container tracing={5}>
        <Row tracing={5}>
          <Column tracing={5} />
        </Row>
        <Row tracing={5}>
          <Column tracing={5} size={6} />
          <Column tracing={5} size={6} />
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
  .add('Form', () => <Form onChange={action('change')} />)
  .add('Input', () => <Input onChange={action('change')} block />)
  .add('DropDown', () => <Input onChange={action('change')} dark />);

storiesOf('Photo Galleries', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Masonry', () => (
    <Masonry
      itemsPerRow={[2, 3]}
      images={[
        { src: '/images/1.jpg', aspectRatio: 3968 / 2976 },
        { src: '/images/2.jpg', aspectRatio: 5344 / 3563 },
        { src: '/images/3.jpg', aspectRatio: 5653 / 3769 },
        { src: '/images/4.jpg', aspectRatio: 3648 / 5472 },
        { src: '/images/5.jpg', aspectRatio: 4570 / 3264 },
        { src: '/images/6.jpg', aspectRatio: 5472 / 3648 },
        { src: '/images/7.jpg', aspectRatio: 122 / 182 },
        { src: '/images/8.jpg', aspectRatio: 122 / 182 }
      ]}
    />
  ));
