import React from 'react';
import { shallow } from 'enzyme';
import MenuBody from '../MenuBody';

it('renders without crashing', () => {
  shallow(<MenuBody />);
});
