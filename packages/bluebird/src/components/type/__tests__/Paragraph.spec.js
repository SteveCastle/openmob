import React from 'react';
import { shallow } from 'enzyme';
import Paragraph from '../Paragraph';

it('renders without crashing', () => {
  shallow(<Paragraph>Hi</Paragraph>);
});
