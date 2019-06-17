import React from 'react';

import { storiesOf } from '@storybook/react';
import App from '../App';
import SideBar from '../sidebar/SideBar';
import Content from '../Content';
import Modal from '../Modal';

import ContentPanel from '../ContentPanel';

storiesOf('Layout System/App', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default App', () => (
    <App>
      <SideBar />
      <ContentPanel>
        <Content>
          <Modal active />
        </Content>
      </ContentPanel>
    </App>
  ));
