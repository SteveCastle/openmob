import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import faker from 'faker';
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
import SubHeader from '../components/type/SubHeader';

storiesOf('Typography', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Content Type Stack', () => (
    <Layout>
      <Container>
        <Header>{faker.name.findName()}</Header>
        <SubHeader>{faker.company.catchPhrase()}</SubHeader>
        <Paragraph>{faker.lorem.paragraph()}</Paragraph>
      </Container>
    </Layout>
  ));

storiesOf('Layout', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Reponsive Flexbox Grid', () => (
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
  .add('Flexbox Container', () => (
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
  .add('Icon Buttons', () => (
    <Button onClick={action('clicked')} label="Hello Button" />
  ))
  .add('Block mode', () => (
    <Button onClick={action('clicked')} label="Hello Button" block />
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
