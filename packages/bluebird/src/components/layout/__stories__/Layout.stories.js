import React from 'react';
import { storiesOf } from '@storybook/react';
import { Layout, Row, Column, Container } from '../';
storiesOf('Layout System/Layout', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Responsive Flexbox Grid', () => (
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
  .add('Row With a Container', () => (
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
